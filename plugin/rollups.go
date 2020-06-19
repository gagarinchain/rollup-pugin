package plugin

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	common_pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/common/rpc"
	"github.com/gagarinchain/rollup-plugin/eth"
	pb "github.com/gagarinchain/rollup-plugin/protobuff"
	"github.com/golang/protobuf/proto"
	"github.com/op/go-logging"
	"google.golang.org/grpc"
	"math/big"
	"net"
	"sort"
)

var log = logging.MustGetLogger("rollups")

type Settings struct {
	Rollups struct {
		ContractAddress   string `yaml:"ContractAddress"`
		BlockchainService string `yaml:"BlockchainService"`
		URL               string `yaml:"URL"`
		PrivateKey        string `yaml:"PrivateKey"`
		Network           struct {
			Address              string `yaml:"Address"`
			MaxConcurrentStreams uint32 `yaml:"MaxConcurrentStreams"`
		} `yaml:"Rpc"`
		Server struct {
			Address              string `yaml:"Address"`
			MaxConcurrentStreams uint32 `yaml:"MaxConcurrentStreams"`
		} `yaml:"Rpc"`
	} `yaml:"Rollups"`
}

type RollupsPlugin struct {
	common_pb.OnBlockCommitServer
	common_pb.OnReceiveProposalServer
	common_pb.OnNewBlockCreatedServer
	client    *rpc.CommonClient
	txOpts    *bind.TransactOpts
	eth       *eth.Eth
	ethClient *ethclient.Client
	address   common.Address
}

type Config struct {
	Address              string
	MaxConcurrentStreams uint32
}

type ByHeight []*common_pb.BlockS

func (h ByHeight) Len() int { return len(h) }
func (h ByHeight) Less(i, j int) bool {
	return h[i].GetHeader().GetHeight() < h[j].GetHeader().GetHeight()
}
func (h ByHeight) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (r *RollupsPlugin) Bootstrap(cfg Config) error {
	log.Info("Bootstrapping rollups plugin")
	lis, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{grpc.MaxConcurrentStreams(cfg.MaxConcurrentStreams)}
	grpcServer := grpc.NewServer(opts...)
	common_pb.RegisterOnBlockCommitServer(grpcServer, r)
	common_pb.RegisterOnReceiveProposalServer(grpcServer, r)
	common_pb.RegisterOnNewBlockCreatedServer(grpcServer, r)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}
	log.Info("Rollups plugin bootstrapped")
	return nil
}

func NewRollupsPlugin(s *Settings) *RollupsPlugin {
	txOpts, instance, ethClient, addr := InitEth(s)
	client := rpc.InitCommonClient(s.Rollups.Network.Address)
	return &RollupsPlugin{
		client:    client,
		txOpts:    txOpts,
		eth:       instance,
		ethClient: ethClient,
		address:   addr,
	}
}

func (r *RollupsPlugin) Close() {
	r.ethClient.Close()
}

func (r *RollupsPlugin) refreshNonce() {
	nonce, err := r.ethClient.PendingNonceAt(context.Background(), r.address)
	if err != nil {
		log.Fatal(err)
	}
	r.txOpts.Nonce = big.NewInt(int64(nonce))
}

//TODO I must be proposer here, validate it
//When block is committed, we send it's rollup to ethereum
//1. Query top height stored on Ethereum
//2. Load fork between topHeight and commitedHeight
//3. Send all rollups to ethereum
func (r *RollupsPlugin) OnBlockCommit(ctx context.Context, req *common_pb.OnBlockCommitRequest) (*common_pb.OnBlockCommitResponse, error) {
	log.Debug("On block commit")

	proposer, err := r.client.Pbc().GetProposerForView(ctx, &common_pb.GetProposerForViewRequest{
		View: req.GetBlock().GetHeader().Height,
	})
	if err != nil {
		log.Error("GetProposerForView failed", err)
		return nil, err
	}

	if !bytes.Equal(proposer.Peer.Address, req.Me.Address) {
		log.Info("we are not proposers, skipping")
		return &common_pb.OnBlockCommitResponse{}, nil
	}

	contractHeight, err := r.eth.GetTopHeight(&bind.CallOpts{})
	if err != nil {
		log.Error("GetTopHeight failed", err)
		return nil, err
	}

	if contractHeight > req.Block.Header.Height {
		log.Error("Ethereum rollups are further than gagarin chain")
		return &common_pb.OnBlockCommitResponse{}, nil
	}

	var fork []*common_pb.BlockS
	if req.Block.Header.Height > contractHeight+1 {
		log.Infof("Loading fork starting at %v to %v", contractHeight, req.Block.Header.Height)
		loaded, err := r.client.Pbc().GetFork(ctx, &common_pb.GetForkRequest{
			Height:   contractHeight + 1,
			HeadHash: req.Block.Header.ParentHash,
		})
		if err != nil {
			log.Error("Get fork failed", err)
			return nil, err
		}

		sort.Sort(ByHeight(loaded.GetBlocks()))

		fork = append(loaded.Blocks, req.Block)
	}

	for _, b := range fork {
		pbBlock := &pb.BlockHeader{
			Hash:       b.Header.Hash,
			ParentHash: b.Header.ParentHash,
			QcHash:     b.Header.QcHash,
			DataHash:   b.Header.DataHash,
			TxHash:     b.Header.TxHash,
			StateHash:  b.Header.StateHash,
			Height:     b.Header.Height,
			Timestamp:  b.Header.Timestamp,
		}

		h, err := proto.Marshal(pbBlock)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		//spew.Dump("Sending header", pbBlock)
		//spew.Dump("Sending rollup", emptyBytes(b.Data.Data))
		log.Infof("Sending rollup at height %v", b.Header.Height)
		r.refreshNonce()
		_, err = r.eth.AddBlock(r.txOpts, h, emptyBytes(b.Data.Data))
		if err != nil {
			log.Error("AddBlock failed", err)
			return nil, err
		}
	}
	log.Infof("%v sent %v rollups", common.Bytes2Hex(req.Me.GetAddress()), len(fork))
	return &common_pb.OnBlockCommitResponse{}, nil
}

func emptyBytes(data []byte) []byte {
	if data == nil {
		return make([]byte, 0)
	}
	return data
}

func (r *RollupsPlugin) BeforeProposedBlockAdded(ctx context.Context,
	in *common_pb.BeforeProposedBlockAddedRequest) (*common_pb.BeforeProposedBlockAddedResponse, error) {
	log.Debugf("Before proposed block %v added", common.Bytes2Hex(in.Proposal.Block.Header.Hash))
	return &common_pb.BeforeProposedBlockAddedResponse{
		Proposal: in.Proposal,
	}, nil
}

//we check received block's data for valid rollups
func (r *RollupsPlugin) AfterProposedBlockAdded(ctx context.Context, in *common_pb.AfterProposedBlockAddedRequest) (*common_pb.AfterProposedBlockAddedResponse, error) {
	log.Debugf("After proposed block %v added", common.Bytes2Hex(in.Proposal.Block.Header.Hash))
	mapping, transactions := r.createRollups(in.Receipts)

	rm := &pb.Rollup{}
	if err := proto.Unmarshal(in.GetProposal().GetBlock().GetData().GetData(), rm); err != nil {
		log.Error("Bad received rollup format", err)
		return nil, err
	}
	if len(mapping) != len(rm.Accounts) || len(transactions) != len(rm.Transactions) {
		err := errors.New("lengths of rollups are not equal")
		log.Error(err)
		return nil, err
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
			return nil, errors.New("mappings are not the same")
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
			return nil, errors.New("bad rollup")
		}
	}

	log.Infof("Valid rollup for %v block", common.Bytes2Hex(in.Proposal.Block.Header.Hash))
	return &common_pb.AfterProposedBlockAddedResponse{}, nil
}

func (r *RollupsPlugin) BeforeVoted(ctx context.Context, in *common_pb.BeforeVotedRequest) (*common_pb.BeforeVotedResponse, error) {
	return &common_pb.BeforeVotedResponse{
		Vote: in.Vote,
	}, nil
}

func (r *RollupsPlugin) AfterVoted(ctx context.Context, in *common_pb.AfterVotedRequest) (*common_pb.AfterVotedResponse, error) {
	return &common_pb.AfterVotedResponse{}, nil
}

//When new block created make rollup using receipts
func (r *RollupsPlugin) OnNewBlockCreated(ctx context.Context, in *common_pb.OnNewBlockCreatedRequest) (*common_pb.OnNewBlockCreatedResponse, error) {
	log.Debugf("Creating rollup for block %v", common.Bytes2Hex(in.Block.Header.Hash))
	if len(in.Receipts) == 0 {
		log.Infof("Block %v has no receipts", common.Bytes2Hex(in.Block.Header.Hash))
		return &common_pb.OnNewBlockCreatedResponse{
			Block: in.Block,
		}, nil
	}

	addresses, txs := r.createRollups(in.Receipts)
	rollup := &pb.Rollup{
		Accounts:     addresses,
		Transactions: txs,
	}

	log.Infof("Created rollup for %v block", common.Bytes2Hex(in.Block.Header.Hash))
	spew.Dump(rollup)

	marshal, err := proto.Marshal(rollup)
	if err != nil {
		return nil, err
	}

	in.Block.Data.Data = marshal
	return &common_pb.OnNewBlockCreatedResponse{
		Block: in.Block,
	}, nil
}

func (r *RollupsPlugin) createRollups(receipts []*common_pb.Receipt) ([][]byte, []*pb.Transaction) {
	spew.Dump("Creating rollup for", receipts)
	var addresses [][]byte
	var txs []*pb.Transaction
	for _, receipt := range receipts {
		//add addresses to map
		from, to := int32(-1), int32(-1)
		for i, v := range addresses {
			if bytes.Equal(v, receipt.GetFrom()) {
				from = int32(i)
			}
			if bytes.Equal(v, receipt.GetTo()) {
				to = int32(i)
			}
			if from > -1 && to > -1 {
				break
			}
		}
		if from == -1 && !bytes.Equal(receipt.GetFrom(), common.Address{}.Bytes()) {
			addresses = append(addresses, receipt.GetFrom())
			from = int32(len(addresses) - 1)
		}
		if to == -1 && !bytes.Equal(receipt.GetTo(), common.Address{}.Bytes()) {
			addresses = append(addresses, receipt.GetTo())
			to = int32(len(addresses) - 1)
		}
		tx := &pb.Transaction{
			From:  from,
			To:    to,
			Value: receipt.GetValue(),
		}
		txs = append(txs, tx)

	}
	return addresses, txs
}

func InitEth(s *Settings) (*bind.TransactOpts, *eth.Eth, *ethclient.Client, common.Address) {
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
	log.Infof("My address %v", fromAddress.Hex())

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
	log.Info("Eth init completed")
	return auth, instance, client, fromAddress
}
