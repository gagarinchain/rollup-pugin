package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagarinchain/common/api"
	common2 "github.com/gagarinchain/common/eth/common"
	pb "github.com/gagarinchain/rollups/protobuff"
	"github.com/golang/protobuf/proto"
	"math/big"
	"time"
)

func main() {
	auth, instance := initEth(ReadSettings())

	account1 := common2.HexToAddress("0x331BcCEb099D3A66e1921C4e434FCBeF853e2B30")
	account2 := common2.HexToAddress("0xd166319B7eBa5DA39882b2B6D17E907801a77234")
	account3 := common2.HexToAddress("0x4595031751B620179FFC09f6E9DB7a5016B53ee4")
	account4 := common2.HexToAddress("0x41E46d84c206007982F93C0874152B8cF6127985")

	receipt1 := &mockReceipt{
		from:  common2.Address{},
		to:    account1,
		value: big.NewInt(100),
	}
	receipt2 := &mockReceipt{
		from:  common2.Address{},
		to:    account2,
		value: big.NewInt(200),
	}
	receipt3 := &mockReceipt{
		from:  common2.Address{},
		to:    account3,
		value: big.NewInt(300),
	}
	receipt4 := &mockReceipt{
		from:  common2.Address{},
		to:    account4,
		value: big.NewInt(400),
	}

	receipt5 := &mockReceipt{
		from:  account1,
		to:    account2,
		value: big.NewInt(10),
	}
	receipt6 := &mockReceipt{
		from:  account2,
		to:    account3,
		value: big.NewInt(20),
	}
	receipt7 := &mockReceipt{
		from:  account3,
		to:    account4,
		value: big.NewInt(30),
	}

	receipt8 := &mockReceipt{
		from:  account2,
		to:    common2.Address{},
		value: big.NewInt(50),
	}
	receipt9 := &mockReceipt{
		from:  account3,
		to:    common2.Address{},
		value: big.NewInt(50),
	}

	h1, b1, err := CreateBlockRollup(receipt1, receipt2, receipt3, receipt4)
	h2, b2, err := CreateBlockRollup(receipt5, receipt6, receipt7)
	h3, b3, err := CreateBlockRollup(receipt8, receipt9)

	h := []*pb.BlockHeader{h1, h2, h3}
	r := []*pb.Rollup{b1, b2, b3}

	for i := 0; i < 3; i++ {
		header, err := proto.Marshal(h[i])
		if err != nil {
			log.Error(err)
		}
		rollup, err := proto.Marshal(r[i])
		if err != nil {
			log.Error(err)
		}

		log.Debugf("Header hex %v", common2.Bytes2Hex(header))
		log.Debugf("Rollup hex %v", common2.Bytes2Hex(rollup))

		_, err = instance.AddBlock(auth, header, rollup)
		if err != nil {
			//log.Fatal(err)
		}

		//log.Debugf("tx sent: %s %s", tx.Hash().Hex(), common2.Bytes2Hex(tx.Data())) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	}

	result, err := instance.GetTopHeight(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Debug(result) // "bar"
}

func CreateBlockRollup(receipt ...api.Receipt) (*pb.BlockHeader, *pb.Rollup, error) {
	plugin := &rollupsPlugin{}
	a1, t1 := plugin.createRollups(receipt)

	rollup1 := &pb.Rollup{
		Accounts:     a1,
		Transactions: t1,
	}

	data1, err := proto.Marshal(rollup1)
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	h1 := &pb.BlockHeader{
		Hash:       crypto.Keccak256([]byte("Hash")),
		ParentHash: crypto.Keccak256([]byte("ParentHash")),
		QcHash:     crypto.Keccak256([]byte("QcHash")),
		DataHash:   crypto.Keccak256(data1),
		TxHash:     crypto.Keccak256([]byte("TxHash")),
		StateHash:  crypto.Keccak256([]byte("StateHash")),
		Height:     10,
		Timestamp:  time.Now().Unix(),
	}

	return h1, rollup1, err
}

type mockReceipt struct {
	txHash    common2.Hash
	txIndex   int32
	from      common2.Address
	to        common2.Address
	value     *big.Int
	toValue   *big.Int
	fromValue *big.Int
}

func (r *mockReceipt) FromValue() *big.Int {
	return r.fromValue
}

func (r *mockReceipt) ToValue() *big.Int {
	return r.toValue
}

func NewReceipt(txHash common2.Hash, txIndex int32, from common2.Address, to common2.Address, value *big.Int, toValue *big.Int, fromValue *big.Int) *mockReceipt {
	return &mockReceipt{txHash: txHash, txIndex: txIndex, from: from, to: to, value: value, toValue: toValue, fromValue: fromValue}
}

func (r *mockReceipt) Value() *big.Int {
	return r.value
}

func (r *mockReceipt) To() common2.Address {
	return r.to
}

func (r *mockReceipt) From() common2.Address {
	return r.from
}

func (r *mockReceipt) TxIndex() int32 {
	return r.txIndex
}

func (r *mockReceipt) TxHash() common2.Hash {
	return r.txHash
}
