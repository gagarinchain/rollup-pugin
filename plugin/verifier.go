package plugin

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagarinchain/common/api"
	c "github.com/gagarinchain/common/eth/common"
	crypto_g "github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/network/blockchain/tx"
	"github.com/gagarinchain/rollup-plugin/gateway"
	"math/big"
)

type Verifier struct {
	gatewayApi        *gateway.Gateway
	ethClient         *ethclient.Client
	txSend            *tx.TxSend
	pk                *crypto_g.PrivateKey
	confirmations     int32
	pendingWaitBlocks int32
	startHeights      map[common.Hash]int64 //we store cache of height for each transaction and keep each record for pendingWaitBlocks

}

func NewVerifier(gatewayApi *gateway.Gateway, ethClient *ethclient.Client, txSend *tx.TxSend, confirmations int32,
	pendingWaitBlocks int32, pk *crypto_g.PrivateKey) *Verifier {
	return &Verifier{
		gatewayApi:        gatewayApi,
		ethClient:         ethClient,
		txSend:            txSend,
		confirmations:     confirmations,
		pendingWaitBlocks: pendingWaitBlocks,
		pk:                pk,
	}

}

func (a *Verifier) verify(ctx context.Context, tpb *pb.TransactionS) (bool, error) {
	pending, err := a.gatewayApi.GetPending(&bind.CallOpts{}, common.BytesToAddress(tpb.From))
	if err != nil {
		return false, err
	}
	pendingHeight := pending.BlockNumber

	latest, err := a.ethClient.HeaderByNumber(ctx, nil)
	if err != nil {
		return false, err
	}

	startHeight, f := a.startHeights[common.BytesToHash(tpb.Hash)]
	if !f {
		a.startHeights[common.BytesToHash(tpb.Hash)] = pendingHeight.Int64()
		startHeight = pendingHeight.Int64()
	}

	if pendingHeight.Int64() == 0 { //no pending yet
		if pendingHeight.Int64()-startHeight > int64(a.pendingWaitBlocks) { //no pending transaction for pendingWaitBlocks
			delete(a.startHeights, common.BytesToHash(tpb.Hash)) //clean cache
			return false, errors.New("no pending found")
		}
		return false, nil
	} else { // we have pending record
		if pendingHeight.Int64() != tpb.Value { //check value equality
			delete(a.startHeights, common.BytesToHash(tpb.Hash)) //clean cache
			return false, errors.New("Pending and Settlement amounts mismatch ")
		}
	}

	if pendingHeight.Int64()+int64(a.confirmations) < latest.Number.Int64() { //confirmed
		opts := &tx.TransactOpts{
			From:       c.Address{},
			PrivateKey: a.pk,
			Fee:        big.NewInt(1),
			Context:    ctx,
		}
		if err := a.txSend.Transact(opts, api.Agreement, c.BytesToAddress(tpb.To), big.NewInt(tpb.Value), nil); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}
