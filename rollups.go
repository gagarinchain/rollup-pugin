package rollups

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gagarinchain/common/api"
	"github.com/gagarinchain/rollup-plugin/eth"
	pb "github.com/gagarinchain/rollup-plugin/protobuff/gagarin/rollups/pb"
	"github.com/golang/protobuf/proto"
	"github.com/op/go-logging"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/big"
	"os"
)

var log = logging.MustGetLogger("rollups")

func InitPlugin(settingsPath string) *rollupsPlugin {
	settings := ReadSettings(settingsPath)
	opts, e := initEth(settings)

	return &rollupsPlugin{
		txOpts: opts,
		eth:    e,
	}
}

type Settings struct {
	Rollups struct {
		ContractAddress string `yaml:"ContractAddress"`
		URL             string `yaml:"URL"`
		PrivateKey      string `yaml:"PrivateKey"`
	} `yaml:"Rollups"`
}

type rollupsPlugin struct {
	settings *Settings
	txOpts   *bind.TransactOpts
	eth      *eth.Eth
}

func NewRollupsPlugin(settings *Settings, txOpts *bind.TransactOpts, eth *eth.Eth) *rollupsPlugin {
	return &rollupsPlugin{settings: settings, txOpts: txOpts, eth: eth}
}

func (r *rollupsPlugin) OnBlockCommit(bc api.Blockchain, block api.Block, orphans *treemap.Map) error {
	contractHeight, err := r.eth.GetTopHeight(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fork := bc.GetFork(contractHeight, block.Header().Hash())
	for _, b := range fork {
		pbBlock := &pb.BlockHeader{
			Hash:       b.Header().Hash().Bytes(),
			ParentHash: b.Header().Parent().Bytes(),
			QcHash:     b.Header().Hash().Bytes(),
			DataHash:   b.Header().DataHash().Bytes(),
			TxHash:     b.Header().TxHash().Bytes(),
			StateHash:  b.Header().StateHash().Bytes(),
			Height:     b.Height(),
			Timestamp:  b.Header().Timestamp().Unix(),
		}

		h, err := proto.Marshal(pbBlock)
		if err != nil {
			log.Error(err)
			return err
		}
		_, err = r.eth.AddBlock(r.txOpts, h, b.Data())
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

func (r *rollupsPlugin) BeforeProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal) error {
	return nil
}

//we check received block's data for valid rollups
func (r *rollupsPlugin) AfterProposedBlockAdded(pacer api.Pacer, bc api.Blockchain, proposal api.Proposal, receipts []api.Receipt) error {
	mapping, transactions := r.createRollups(receipts)

	rm := &pb.Rollup{}
	if err := proto.Unmarshal(proposal.NewBlock().Data(), rm); err != nil {
		return err
	}

	if len(mapping) != len(rm.Accounts) || len(transactions) != len(rm.Accounts) {
		return errors.New("lengths of rollups are not equal")
	}

	//we must have exact the same rollups, probably with permutated mappings, so let's transform mappings and validate transactions
	transformation := make(map[int32]int32)
	for i, v := range mapping {
		found := false
		for i1, v1 := range rm.Accounts {
			if bytes.Equal(v, v1) {
				transformation[int32(i)] = int32(i1)
				found = true
				break
			}
		}
		if !found {
			log.Errorf("No address %v:%v is found in mapping", i, v)
			return errors.New("mappings are not the same")
		}
	}
	for _, txs := range transactions {
		found := false
		for _, txs1 := range rm.Transactions {
			if transformation[txs.From] == txs1.From && transformation[txs.To] == txs1.To && txs.Value == txs1.Value {
				found = true
				break
			}
		}
		if !found {
			log.Errorf("not found transaction %v/%v/%v in rollup", txs.From, txs.To, txs.Value)
			return errors.New("bad rollup")
		}
	}

	return nil
}

func (r *rollupsPlugin) BeforeVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	return nil
}

func (r *rollupsPlugin) AfterVoted(pacer api.Pacer, bc api.Blockchain, vote api.Vote) error {
	return nil
}

func (r *rollupsPlugin) OnBlockCreated(bc api.Blockchain, blockBuilder api.BlockBuilder, receipts []api.Receipt) error {
	log.Debug("Creating rollup")
	addresses, txs := r.createRollups(receipts)
	rollup := &pb.Rollup{
		Accounts:     addresses,
		Transactions: txs,
	}

	marshal, err := proto.Marshal(rollup)
	if err != nil {
		return err
	}
	blockBuilder.SetData(marshal)

	return nil
}

func (r *rollupsPlugin) createRollups(receipts []api.Receipt) ([][]byte, []*pb.Transaction) {
	var addresses [][]byte
	var txs []*pb.Transaction
	for _, receipt := range receipts {
		//add addresses to map
		from, to := int32(-1), int32(-1)
		for i, v := range addresses {
			if bytes.Equal(v, receipt.From().Bytes()) {
				from = int32(i)
			}
			if bytes.Equal(v, receipt.To().Bytes()) {
				to = int32(i)
			}
			if from > -1 && to > -1 {
				break
			}
		}
		if from == -1 && !bytes.Equal(receipt.From().Bytes(), common.Address{}.Bytes()) {
			addresses = append(addresses, receipt.From().Bytes())
			from = int32(len(addresses) - 1)
		}
		if to == -1 && !bytes.Equal(receipt.To().Bytes(), common.Address{}.Bytes()) {
			addresses = append(addresses, receipt.To().Bytes())
			to = int32(len(addresses) - 1)
		}
		tx := &pb.Transaction{
			From:  from,
			To:    to,
			Value: receipt.Value().Int64(),
		}
		txs = append(txs, tx)

	}
	return addresses, txs
}

func initEth(s *Settings) (*bind.TransactOpts, *eth.Eth) {
	client, err := ethclient.Dial(s.Rollups.URL)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(s.Rollups.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(s.Rollups.ContractAddress)
	log.Debug(address.Hex())
	instance, err := eth.NewEth(address, client)
	if err != nil {
		log.Fatal(err)
	}
	return auth, instance
}

func ReadSettings(settingsPath string) (s *Settings) {
	file, e := os.Open(settingsPath)
	if e != nil {
		log.Error("Can't load settings, using default", e)
	} else {
		defer file.Close()
		s = &Settings{}
		byteValue, _ := ioutil.ReadAll(file)
		if err := yaml.Unmarshal(byteValue, s); err != nil {
			log.Error("Can't load settings, using default", e)
		}
	}
	if s == nil {
		s = &Settings{
			Rollups: struct {
				ContractAddress string `yaml:"ContractAddress"`
				URL             string `yaml:"URL"`
				PrivateKey      string `yaml:"PrivateKey"`
			}{
				ContractAddress: "",
				URL:             "",
				PrivateKey:      "",
			}}
	}
	return s
}
