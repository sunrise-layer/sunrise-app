package testnode

// import (
// 	"time"

// 	"github.com/sunriselayer/sunrise/cmd/sunrised/cmd"
// 	"github.com/sunriselayer/sunrise/pkg/appconsts"
// 	"github.com/sunriselayer/sunrise/test/util/genesis"

// 	tmconfig "github.com/cometbft/cometbft/config"
// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	tmtypes "github.com/cometbft/cometbft/types"
// 	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
// 	srvtypes "github.com/cosmos/cosmos-sdk/server/types"
// )

// const (
// 	kibibyte                    = 1024      // bytes
// 	mebibyte                    = 1_048_576 // bytes
// 	DefaultValidatorAccountName = "validator"
// )

// type UniversalTestingConfig struct {
// 	// TmConfig is the Tendermint configuration used for the network.
// 	TmConfig *tmconfig.Config
// 	// AppConfig is the application configuration of the test node.
// 	AppConfig *srvconfig.Config
// 	// AppOptions are the application options of the test node.
// 	AppOptions *KVAppOptions
// 	// AppCreator is used to create the application for the testnode.
// 	AppCreator srvtypes.AppCreator
// 	// SuppressLogs in testnode. This should be set to true when running
// 	// testground tests.
// 	SuppressLogs bool
// }

// // Config is the configuration of a test node.
// type Config struct {
// 	Genesis *genesis.Genesis
// 	UniversalTestingConfig
// }

// func (c *Config) WithGenesis(g *genesis.Genesis) *Config {
// 	c.Genesis = g
// 	return c
// }

// // WithTendermintConfig sets the TmConfig and returns the *Config.
// func (c *Config) WithTendermintConfig(conf *tmconfig.Config) *Config {
// 	c.TmConfig = conf
// 	return c
// }

// // WithAppConfig sets the AppConfig and returns the Config.
// func (c *Config) WithAppConfig(conf *srvconfig.Config) *Config {
// 	c.AppConfig = conf
// 	return c
// }

// // WithAppOptions sets the AppOptions and returns the Config.
// func (c *Config) WithAppOptions(opts *KVAppOptions) *Config {
// 	c.AppOptions = opts
// 	return c
// }

// // WithAppCreator sets the AppCreator and returns the Config.
// func (c *Config) WithAppCreator(creator srvtypes.AppCreator) *Config {
// 	c.AppCreator = creator
// 	return c
// }

// // WithSuppressLogs sets the SuppressLogs and returns the Config.
// func (c *Config) WithSuppressLogs(sl bool) *Config {
// 	c.SuppressLogs = sl
// 	return c
// }

// // WithTimeoutCommit sets the TimeoutCommit and returns the Config.
// func (c *Config) WithTimeoutCommit(d time.Duration) *Config {
// 	c.TmConfig.Consensus.TimeoutCommit = d
// 	return c
// }

// // WithFundedAccounts sets the genesis accounts and returns the Config.
// func (c *Config) WithFundedAccounts(accounts ...string) *Config {
// 	c.Genesis = c.Genesis.WithAccounts(
// 		genesis.NewAccounts(999999999999999999, accounts...)...,
// 	)
// 	return c
// }

// // WithModifiers sets the genesis options and returns the Config.
// func (c *Config) WithModifiers(ops ...genesis.Modifier) *Config {
// 	c.Genesis = c.Genesis.WithModifiers(ops...)
// 	return c
// }

// // WithGenesisTime sets the genesis time and returns the Config.
// func (c *Config) WithGenesisTime(t time.Time) *Config {
// 	c.Genesis = c.Genesis.WithGenesisTime(t)
// 	return c
// }

// // WithChainID sets the chain ID and returns the Config.
// func (c *Config) WithChainID(id string) *Config {
// 	c.Genesis = c.Genesis.WithChainID(id)
// 	return c
// }

// // WithConsensusParams sets the consensus params and returns the Config.
// func (c *Config) WithConsensusParams(params *tmproto.ConsensusParams) *Config {
// 	c.Genesis = c.Genesis.WithConsensusParams(params)
// 	return c
// }

// // DefaultConfig returns the default configuration of a test node.
// func DefaultConfig() *Config {
// 	cfg := &Config{}
// 	return cfg.
// 		WithGenesis(
// 			genesis.NewDefaultGenesis().
// 				WithValidators(genesis.NewDefaultValidator(DefaultValidatorAccountName)),
// 		).
// 		WithTendermintConfig(DefaultTendermintConfig()).
// 		WithAppConfig(DefaultAppConfig()).
// 		WithAppOptions(DefaultAppOptions()).
// 		WithAppCreator(cmd.NewAppServer).
// 		WithSuppressLogs(true).
// 		WithConsensusParams(DefaultConsensusParams())
// }

// func DefaultConsensusParams() *tmproto.ConsensusParams {
// 	cparams := tmtypes.DefaultConsensusParams().ToProto()
// 	cparams.Block.MaxBytes = appconsts.DefaultMaxBytes
// 	cparams.Version.App = appconsts.LatestVersion
// 	return &cparams
// }

// func DefaultTendermintConfig() *tmconfig.Config {
// 	tmCfg := tmconfig.DefaultConfig()
// 	// Reduce the timeout commit to 1ms to speed up the rate at which the test
// 	// node produces blocks.
// 	tmCfg.Consensus.TimeoutCommit = 1 * time.Millisecond

// 	// Override the mempool's MaxTxBytes to allow the testnode to accept a
// 	// transaction that fills the entire square. Any blob transaction larger
// 	// than the square size will still fail no matter what.
// 	maxTxBytes := appconsts.DefaultSquareSizeUpperBound * appconsts.DefaultSquareSizeUpperBound * appconsts.ContinuationSparseShareContentSize
// 	tmCfg.Mempool.MaxTxBytes = maxTxBytes

// 	// Override the MaxBodyBytes to allow the testnode to accept very large
// 	// transactions and respond to queries with large responses (200 MiB was
// 	// chosen only as an arbitrary large number).
// 	tmCfg.RPC.MaxBodyBytes = 200 * mebibyte

// 	return tmCfg
// }
