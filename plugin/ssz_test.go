package plugin

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagarinchain/common"
	"github.com/gagarinchain/common/api"
	c "github.com/gagarinchain/common/eth/common"
	crypto_g "github.com/gagarinchain/common/eth/crypto"
	"github.com/gagarinchain/network/blockchain/tx"
	"github.com/prysmaticlabs/go-ssz"
	"math/big"
	"testing"
	"time"
)

func TestSszSimple(t *testing.T) {
	type T struct {
		From  int32
		To    int32
		Value uint64
		//non serialized field for map

	}
	type P struct {
		Hash      []byte
		Height    uint32
		Timestamp uint64
		Txs       []*T
		Data      []byte
	}

	p := P{
		Hash:      []byte{0x33, 0x34, 0x35, 0x36, 0x34, 0x35, 0x36},
		Height:    1,
		Timestamp: uint64(time.Now().Unix()),
		Data:      []byte{0x11, 0x12, 0x13, 0x14, 0x15},
		Txs: []*T{{
			From:  -1,
			To:    -1,
			Value: 10,
		}, {
			From:  10,
			To:    12,
			Value: 4,
		}, {
			From:  4,
			To:    43,
			Value: 33,
		}},
	}

	marshal, err := ssz.Marshal(p)

	spew.Dump(err)
	spew.Dump(marshal)
}

type Rollup struct {
	Accounts     [][20]byte
	Transactions []*Transaction
}

type Transaction struct {
	From  int32
	To    int32
	Value uint64
}

type SendableHeader struct {
	Height    uint32
	Hash      [32]byte
	TxHash    [32]byte
	StateHash [32]byte
	DataHash  [32]byte
	QcHash    [32]byte
	Parent    [32]byte
	Timestamp uint64
}

func TestSszRollup(t *testing.T) {
	a1 := common.GenerateAddress()
	a2 := common.GenerateAddress()
	a3 := common.GenerateAddress()
	rollup := Rollup{
		Accounts: [][20]byte{a1, a2, a3},
		Transactions: []*Transaction{
			{
				From:  1,
				To:    3,
				Value: 10,
			},
			{
				From:  4,
				To:    5,
				Value: 33,
			},
		},
	}

	p := new([32]byte)

	h := &SendableHeader{
		Height:    15,
		Hash:      crypto.Keccak256Hash([]byte("Hash")),
		TxHash:    crypto.Keccak256Hash([]byte("TxHash")),
		StateHash: *p,
		DataHash:  crypto.Keccak256Hash([]byte("DataHash")),
		QcHash:    crypto.Keccak256Hash([]byte("QcHash")),
		Parent:    crypto.Keccak256Hash([]byte("Parent")),
		Timestamp: uint64(time.Now().Unix()),
	}

	marshal, err := ssz.Marshal(rollup)
	mh, err := ssz.Marshal(h)

	spew.Dump(err)
	spew.Dump(a1)
	spew.Dump(a2)
	spew.Dump(a3)
	spew.Dump(marshal)
	spew.Dump(h)
	spew.Dump(mh)
	sprintf := fmt.Sprintf("%x", marshal)
	println(sprintf)
	sh := fmt.Sprintf("%x", mh)
	println(sh)
}

func TestRollupGenerate(t *testing.T) {
	a0 := c.HexToAddress("0x41E46d84c206007982F93C0874152B8cF6127985")
	a1 := c.HexToAddress("0xDd9811Cfc24aB8d56036A8ecA90C7B8C75e35950")
	a2 := c.HexToAddress("0xf3AA514423aE2c6f66497D69f2fc899f0ad25b00")
	a3 := c.HexToAddress("0x95AC9774Cf32AdB66f76726502Fc0C223bCE13e2")
	rollup := Rollup{
		Accounts: [][20]byte{a0, a1, a2, a3},
		Transactions: []*Transaction{
			{
				From:  1,
				To:    2,
				Value: 5,
			}, {
				From:  1,
				To:    3,
				Value: 4,
			},
			{
				From:  0,
				To:    1,
				Value: 13,
			}, {
				From:  2,
				To:    0,
				Value: 8,
			}, {
				From:  0,
				To:    2,
				Value: 16,
			},
		},
	}

	marshal, _ := ssz.Marshal(rollup)
	sprint := fmt.Sprintf("%x", marshal)
	println(sprint)

	h := &SendableHeader{
		Height:    1,
		Hash:      crypto.Keccak256Hash([]byte("Hash")),
		TxHash:    crypto.Keccak256Hash([]byte("TxHash")),
		StateHash: *new([32]byte),
		DataHash:  crypto.Keccak256Hash(marshal),
		QcHash:    crypto.Keccak256Hash([]byte("QcHash")),
		Parent:    crypto.Keccak256Hash([]byte("Parent")),
		Timestamp: uint64(time.Now().Unix()),
	}

	mh, _ := ssz.Marshal(h)

	s := fmt.Sprintf("%x", mh)
	println(s)
}

func TestTxSend(t *testing.T) {
	send := tx.CreateTxSend("localhost:9280")

	if err := send.Start(); err != nil {
		t.Error(err)
		return
	}

	b := c.Hex2Bytes("4a157ef4295be2413c59613cd5dd4f3b0688ac6227709a9634608f33a49250ef")
	pk := crypto_g.PkFromBytes(b)

	timeout, _ := context.WithTimeout(context.Background(), time.Second)
	send.Transact(&tx.TransactOpts{
		From:       c.HexToAddress("0xdd9811cfc24ab8d56036a8eca90c7b8c75e35950"),
		PrivateKey: pk,
		Nonce:      uint64(2),
		Fee:        big.NewInt(1),
		Context:    timeout,
	}, api.Payment, c.HexToAddress("0x3db5e6f31B6b46ffeC25E512763FC61aEc42654c"), big.NewInt(10), nil)
}
