package keeper

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/oraichain/orai/packages/filecache"
	"github.com/oraichain/orai/x/provider/types"
)

// Implements OracleScriptSet interface
var _ types.OracleScriptSet = Keeper{}

// Implements AIDataSourceSet interface
var _ types.AIDataSourceSet = Keeper{}

// Implements TestCaseSet interface
var _ types.TestCaseSet = Keeper{}

// Keeper of the provider store
type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              *codec.Codec
	paramSpace       params.Subspace
	supplyKeeper     types.SupplyKeeper
	bankKeeper       types.BankKeeper
	stakingKeeper    types.StakingKeeper
	DistrKeeper      types.DistrKeeper
	fileCache        filecache.Cache
	feeCollectorName string
}

// NewKeeper creates a provider keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, providerSubspace params.Subspace, supplyKeeper types.SupplyKeeper, bankKeeper types.BankKeeper, stakingKeeper types.StakingKeeper, distrKeeper types.DistrKeeper, feeCollectorName string, cacheDir string) Keeper {
	if !providerSubspace.HasKeyTable() {
		// register parameters of the provider module into the param space
		providerSubspace = providerSubspace.WithKeyTable(types.ParamKeyTable())
	}
	fileCache := filecache.New(cacheDir)
	return Keeper{
		storeKey:         key,
		cdc:              cdc,
		paramSpace:       providerSubspace,
		fileCache:        fileCache,
		supplyKeeper:     supplyKeeper,
		bankKeeper:       bankKeeper,
		stakingKeeper:    stakingKeeper,
		DistrKeeper:      distrKeeper,
		feeCollectorName: feeCollectorName,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

//IsNamePresent checks if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

// // CreateStrategy allows users to create a new strategy into the store
// func (k Keeper) CreateStrategy(ctx sdk.Context, name string, strategy types.Strategy) {
// 	store := ctx.KVStore(k.storeKey)

// 	bz := k.cdc.MustMarshalBinaryLengthPrefixed(strategy)
// 	store.Set(types.StrategyStoreKey(strategy.StratID, name), bz)
// }

// GetParam returns the parameter as specified by key as an uint64.
func (k Keeper) GetParam(ctx sdk.Context, key []byte) (res uint64) {
	k.paramSpace.Get(ctx, key, &res)
	return res
}

// SetParam saves the given key-value parameter to the store.
func (k Keeper) SetParam(ctx sdk.Context, key []byte, value uint64) {
	k.paramSpace.Set(ctx, key, value)
}

// GetParams returns all current parameters as a types.Params instance.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// GetDNamesTcNames is a function that collects test case names and data source names from the oracle script
func (k Keeper) GetDNamesTcNames(oscriptPath string) ([]string, []string, error) {
	//use "data source" as an argument to collect the data source script name
	cmd := exec.Command("bash", oscriptPath, "aiDataSource")
	cmd.Stdin = strings.NewReader("some input")
	var dataSourceName bytes.Buffer
	cmd.Stdout = &dataSourceName
	err := cmd.Run()
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrFailedToOpenFile, err.Error())
	}

	// collect data source result from the script
	result := strings.TrimSuffix(dataSourceName.String(), "\n")

	aiDataSources := strings.Fields(result)

	//use "test case" as an argument to collect the test case script name
	cmd = exec.Command("bash", oscriptPath, "testcase")
	cmd.Stdin = strings.NewReader("some input")
	var testCaseName bytes.Buffer
	cmd.Stdout = &testCaseName
	err = cmd.Run()
	if err != nil {
		return nil, nil, sdkerrors.Wrap(types.ErrFailedToOpenFile, fmt.Sprintf("failed to collect test case name: %s", result))
	}

	// collect data source result from the script
	result = strings.TrimSuffix(testCaseName.String(), "\n")

	testCases := strings.Fields(result)

	return aiDataSources, testCases, nil
}

// GetMinimumFees collects minimum fees needed of an oracle script
func (k Keeper) GetMinimumFees(ctx sdk.Context, dNames, tcNames []string, valNum int) (sdk.Coins, error) {
	var totalFees sdk.Coins
	// we have different test cases, so we need to loop through them
	for i := 0; i < len(tcNames); i++ {
		// loop to run the test case
		// collect all the test cases object to store in the ai request
		testCase, err := k.GetTestCase(ctx, tcNames[i])
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrTestCaseNotFound, fmt.Sprintf("failed to get test case: %s", err.Error()))
		}
		// Aggregate the required fees for an AI request
		totalFees = totalFees.Add(testCase.GetFees()...)
	}

	for j := 0; j < len(dNames); j++ {
		// collect all the data source objects to store in the ai request
		aiDataSource, err := k.GetAIDataSource(ctx, dNames[j])
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrDataSourceNotFound, fmt.Sprintf("failed to get data source: %s", err.Error()))
		}
		// Aggregate the required fees for an AI request
		totalFees = totalFees.Add(aiDataSource.GetFees()...)
	}
	rewardRatio := sdk.NewDecWithPrec(int64(k.GetParam(ctx, types.KeyOracleScriptRewardPercentage)), 2)
	//valFees = 2/5 total dsource and test case fees (70% total in 100% of total fees split into 20% and 50% respectively)
	valFees, _ := sdk.NewDecCoinsFromCoins(totalFees...).MulDec(sdk.NewDecWithPrec(int64(40), 2)).TruncateDecimal()
	//50% + 20% = 70% * validatorCount fees (since k validators will execute, the fees need to be propotional to the number of vals)
	bothFees := sdk.NewDecCoinsFromCoins(totalFees.Add(valFees...)...)
	finalFees, _ := bothFees.MulDec(sdk.NewDec(int64(valNum))).TruncateDecimal()
	minimumFees, _ := sdk.NewDecCoinsFromCoins(finalFees...).QuoDec(rewardRatio).TruncateDecimal()
	fmt.Println("minimum fees: ", minimumFees)
	return minimumFees, nil
}
