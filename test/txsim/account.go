package txsim

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"math"
// 	"strings"
// 	"sync"
// 	"time"

// 	"cosmossdk.io/x/feegrant"
// 	tmservice "github.com/cosmos/cosmos-sdk/client/grpc/cmtservice"
// 	"github.com/cosmos/cosmos-sdk/crypto/hd"
// 	"github.com/cosmos/cosmos-sdk/crypto/keyring"
// 	"github.com/cosmos/cosmos-sdk/types"
// 	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	"github.com/rs/zerolog/log"
// 	"github.com/sunriselayer/sunrise/app/encoding"
// 	"github.com/sunriselayer/sunrise/pkg/appconsts"
// 	"github.com/sunriselayer/sunrise/pkg/user"
// 	"google.golang.org/grpc"
// )

// const defaultFee = DefaultGasLimit * appconsts.DefaultMinGasPrice

// type AccountManager struct {
// 	keys        keyring.Keyring
// 	conn        *grpc.ClientConn
// 	pending     []*account
// 	encCfg      encoding.Config
// 	pollTime    time.Duration
// 	useFeegrant bool

// 	// to protect from concurrent writes to the map
// 	mtx          sync.Mutex
// 	master       *user.Signer
// 	balance      uint64
// 	latestHeight uint64
// 	lastUpdated  time.Time
// 	subaccounts  map[string]*user.Signer
// }

// func NewAccountManager(
// 	ctx context.Context,
// 	keys keyring.Keyring,
// 	encCfg encoding.Config,
// 	masterAccName string,
// 	conn *grpc.ClientConn,
// 	pollTime time.Duration,
// 	useFeegrant bool,
// ) (*AccountManager, error) {
// 	records, err := keys.List()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(records) == 0 {
// 		return nil, fmt.Errorf("no accounts found in keyring")
// 	}

// 	am := &AccountManager{
// 		keys:        keys,
// 		subaccounts: make(map[string]*user.Signer),
// 		encCfg:      encCfg,
// 		pending:     make([]*account, 0),
// 		conn:        conn,
// 		pollTime:    pollTime,
// 		useFeegrant: useFeegrant,
// 	}

// 	if masterAccName == "" {
// 		masterAccName, err = am.findWealthiestAccount(ctx)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if err := am.setupMasterAccount(ctx, masterAccName); err != nil {
// 		return nil, err
// 	}

// 	return am, nil
// }

// func (am *AccountManager) findWealthiestAccount(ctx context.Context) (string, error) {
// 	am.mtx.Lock()
// 	defer am.mtx.Unlock()

// 	records, err := am.keys.List()
// 	if err != nil {
// 		return "", err
// 	}

// 	var (
// 		highestBalance    uint64
// 		wealthiestAddress string
// 	)

// 	for _, record := range records {
// 		address, err := record.GetAddress()
// 		if err != nil {
// 			return "", fmt.Errorf("error getting address for account %s: %w", record.Name, err)
// 		}

// 		if wealthiestAddress == "" {
// 			wealthiestAddress = record.Name
// 			highestBalance = 0
// 		}

// 		// search for the account on chain
// 		balance, err := am.getBalance(ctx, address)
// 		if err != nil {
// 			log.Err(err).Str("account", record.Name).Msg("error getting initial account balance")
// 			continue
// 		}

// 		if balance > highestBalance {
// 			wealthiestAddress = record.Name
// 			highestBalance = balance
// 		}
// 	}

// 	if wealthiestAddress == "" {
// 		return "", fmt.Errorf("no suitable master account found")
// 	}

// 	return wealthiestAddress, nil
// }

// // setupMasterAccount loops through all accounts in the keyring and picks out the one with
// // the highest balance as the master account. Accounts that don't yet exist on chain are
// // ignored.
// func (am *AccountManager) setupMasterAccount(ctx context.Context, masterAccName string) error {
// 	masterRecord, err := am.keys.Key(masterAccName)
// 	if err != nil {
// 		return fmt.Errorf("error getting master account %s: %w", masterAccName, err)
// 	}

// 	masterAddress, err := masterRecord.GetAddress()
// 	if err != nil {
// 		return fmt.Errorf("error getting address for account %s: %w", masterAccName, err)
// 	}

// 	// search for the account on chain
// 	am.balance, err = am.getBalance(ctx, masterAddress)
// 	if err != nil {
// 		return fmt.Errorf("error getting master account %s balance: %w", masterAccName, err)
// 	}

// 	am.master, err = user.SetupSigner(ctx, am.keys, am.conn, masterAddress, am.encCfg)
// 	if err != nil {
// 		return err
// 	}

// 	log.Info().
// 		Str("address", am.master.Address().String()).
// 		Uint64("balance", am.balance).
// 		Msg("set master account")

// 	return nil
// }

// // AllocateAccounts is used by sequences to specify the number of accounts
// // and the balance of each of those accounts. Not concurrently safe.
// func (am *AccountManager) AllocateAccounts(n, balance int) []types.AccAddress {
// 	if n < 1 {
// 		panic("n must be greater than 0")
// 	}
// 	if balance < 1 {
// 		panic("balance must be greater than 0")
// 	}

// 	path := hd.CreateHDPath(types.CoinType, 0, 0).String()
// 	addresses := make([]types.AccAddress, n)
// 	for i := 0; i < n; i++ {
// 		record, _, err := am.keys.NewMnemonic(am.nextAccountName(), keyring.English, path, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
// 		if err != nil {
// 			panic(err)
// 		}
// 		addresses[i], err = record.GetAddress()
// 		if err != nil {
// 			panic(err)
// 		}

// 		am.pending = append(am.pending, &account{
// 			address: addresses[i],
// 			balance: uint64(balance),
// 		})
// 	}
// 	return addresses
// }

// // Submit executes on an operation. This is thread safe.
// func (am *AccountManager) Submit(ctx context.Context, op Operation) error {
// 	if len(op.Msgs) == 0 {
// 		return errors.New("operation must contain at least one message")
// 	}

// 	var address types.AccAddress
// 	for _, msg := range op.Msgs {
// 		signers, _, err := am.encCfg.Codec.GetMsgV1Signers(msg)
// 		if err != nil {
// 			return err
// 		}

// 		if len(signers) != 1 {
// 			return fmt.Errorf("only a single signer is supported got: %d", len(signers))
// 		}

// 		if address == nil {
// 			address = signers[0]
// 		}
// 	}

// 	// If a delay is set, wait for that many blocks to have been produced
// 	// before continuing
// 	if op.Delay != 0 {
// 		if err := am.waitDelay(ctx, uint64(op.Delay)); err != nil {
// 			return fmt.Errorf("error delaying tx submission: %w", err)
// 		}
// 	}

// 	signer, err := am.getSubAccount(address)
// 	if err != nil {
// 		return err
// 	}

// 	opts := make([]user.TxOption, 0)
// 	if op.GasLimit == 0 {
// 		opts = append(opts, user.SetGasLimit(DefaultGasLimit), user.SetFee(defaultFee))
// 	} else {
// 		opts = append(opts, user.SetGasLimit(op.GasLimit))
// 		if op.GasPrice > 0 {
// 			opts = append(opts, user.SetFee(uint64(math.Ceil(float64(op.GasLimit)*op.GasPrice))))
// 		} else {
// 			opts = append(opts, user.SetFee(uint64(math.Ceil(float64(op.GasLimit)*appconsts.DefaultMinGasPrice))))
// 		}
// 	}

// 	if am.useFeegrant {
// 		opts = append(opts, user.SetFeeGranter(am.master.Address()))
// 	}

// 	var res *types.TxResponse
// 	if len(op.Blobs) > 0 {
// 		res, err = signer.SubmitPayForBlob(ctx, op.Blobs, opts...)
// 	} else {
// 		res, err = signer.SubmitTx(ctx, op.Msgs, opts...)
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	// update the latest latestHeight
// 	am.setLatestHeight(res.Height)

// 	log.Info().
// 		Int64("height", res.Height).
// 		Str("address", address.String()).
// 		Str("msgs", msgsToString(op.Msgs)).
// 		Msg("tx committed")

// 	return nil
// }

// // Generate the pending accounts by sending the adequate funds. This operation
// // is not concurrently safe.
// func (am *AccountManager) GenerateAccounts(ctx context.Context) error {
// 	if len(am.pending) == 0 {
// 		return nil
// 	}

// 	msgs := make([]types.Msg, 0)
// 	gasLimit := 0
// 	// batch together all the messages needed to create all the accounts
// 	for _, acc := range am.pending {
// 		if am.balance < acc.balance {
// 			return fmt.Errorf("master account has insufficient funds. has: %v needed: %v", am.balance, acc.balance)
// 		}

// 		if am.useFeegrant {
// 			// create a feegrant message so that the master account pays for all the fees of the sub accounts
// 			feegrantMsg, err := feegrant.NewMsgGrantAllowance(&feegrant.BasicAllowance{}, am.master.Address(), acc.address)
// 			if err != nil {
// 				return fmt.Errorf("error creating feegrant message: %w", err)
// 			}
// 			msgs = append(msgs, feegrantMsg)
// 			gasLimit += FeegrantGasLimit
// 		}

// 		bankMsg := bank.NewMsgSend(am.master.Address(), acc.address, types.NewCoins(types.NewInt64Coin(appconsts.BondDenom, int64(acc.balance))))
// 		msgs = append(msgs, bankMsg)
// 		gasLimit += SendGasLimit
// 	}

// 	err := am.Submit(ctx, Operation{Msgs: msgs, GasLimit: uint64(gasLimit)})
// 	if err != nil {
// 		return fmt.Errorf("error funding accounts: %w", err)
// 	}

// 	// check that the account now exists
// 	for _, acc := range am.pending {
// 		signer, err := user.SetupSigner(ctx, am.keys, am.conn, acc.address, am.encCfg)
// 		if err != nil {
// 			return err
// 		}

// 		signer.SetPollTime(am.pollTime)

// 		// set the account
// 		am.mtx.Lock()
// 		am.subaccounts[acc.address.String()] = signer
// 		am.mtx.Unlock()
// 		log.Info().
// 			Str("address", acc.address.String()).
// 			Uint64("balance", acc.balance).
// 			Uint64("account number", signer.AccountNumber()).
// 			Msg("initialized account")
// 	}

// 	// clear the pending accounts
// 	am.pending = nil
// 	return nil
// }

// // getBalance returns the balance for the given address
// func (am *AccountManager) getBalance(ctx context.Context, address types.AccAddress) (uint64, error) {
// 	balanceResp, err := bank.NewQueryClient(am.conn).Balance(ctx, &bank.QueryBalanceRequest{
// 		Address: address.String(),
// 		Denom:   appconsts.BondDenom,
// 	})
// 	if err != nil {
// 		return 0, fmt.Errorf("error getting balance for %s: %w", address.String(), err)
// 	}
// 	return balanceResp.GetBalance().Amount.Uint64(), nil
// }

// func (am *AccountManager) getSubAccount(address types.AccAddress) (*user.Signer, error) {
// 	am.mtx.Lock()
// 	defer am.mtx.Unlock()
// 	signer, exists := am.subaccounts[address.String()]
// 	if !exists {
// 		if bytes.Equal(am.master.Address(), address) {
// 			return am.master, nil
// 		}
// 		return nil, fmt.Errorf("account %s does not exist", address)
// 	}
// 	return signer, nil
// }

// func (am *AccountManager) waitDelay(ctx context.Context, blocks uint64) error {
// 	latestHeight, err := am.updateHeight(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	ticker := time.NewTicker(am.pollTime)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err()
// 		case <-ticker.C:
// 			height, err := am.updateHeight(ctx)
// 			if err != nil {
// 				return err
// 			}
// 			if height > latestHeight+blocks {
// 				return nil
// 			}
// 		}
// 	}
// }

// func (am *AccountManager) setLatestHeight(height int64) uint64 {
// 	am.mtx.Lock()
// 	defer am.mtx.Unlock()
// 	if uint64(height) > am.latestHeight {
// 		am.latestHeight = uint64(height)
// 		am.lastUpdated = time.Now()
// 	}
// 	return am.latestHeight
// }

// func (am *AccountManager) updateHeight(ctx context.Context) (uint64, error) {
// 	am.mtx.Lock()
// 	if time.Since(am.lastUpdated) < am.pollTime {
// 		am.mtx.Unlock()
// 		return am.latestHeight, nil
// 	}
// 	am.mtx.Unlock()
// 	resp, err := tmservice.NewServiceClient(am.conn).GetLatestBlock(ctx, &tmservice.GetLatestBlockRequest{})
// 	if err != nil {
// 		return 0, err
// 	}
// 	return am.setLatestHeight(resp.SdkBlock.Header.Height), nil
// }

// func (am *AccountManager) nextAccountName() string {
// 	am.mtx.Lock()
// 	defer am.mtx.Unlock()
// 	return accountName(len(am.pending) + len(am.subaccounts))
// }

// type account struct {
// 	address types.AccAddress
// 	balance uint64
// }

// func accountName(n int) string { return fmt.Sprintf("tx-sim-%d", n) }

// func msgsToString(msgs []types.Msg) string {
// 	msgsStr := make([]string, len(msgs))
// 	for i, msg := range msgs {
// 		msgsStr[i] = types.MsgTypeURL(msg)
// 	}
// 	return strings.Join(msgsStr, ",")
// }
