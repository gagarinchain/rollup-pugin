package plugin

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	c2 "github.com/gagarinchain/common"
	"github.com/gagarinchain/common/api"
	c "github.com/gagarinchain/common/eth/common"
	crypto_g "github.com/gagarinchain/common/eth/crypto"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/common/rpc"
	"github.com/gagarinchain/network/blockchain/tx"
	"github.com/golang/protobuf/proto"
	"math/big"
)

type Prover struct {
	txSend    *tx.TxSend
	pk        *crypto_g.PrivateKey
	client    *rpc.CommonClient
	committee []*pb.Peer
	proofs    map[c.Address][][]byte
}

func NewProver(txSend *tx.TxSend, pk *crypto_g.PrivateKey, client *rpc.CommonClient) *Prover {
	return &Prover{txSend: txSend, pk: pk, client: client, proofs: make(map[c.Address][][]byte)}
}

func (p *Prover) Bootstrap(ctx context.Context) error {
	resp, err := p.client.Pbc().GetCommittee(ctx, &pb.GetCommitteeRequest{})
	if err != nil {
		return err
	}

	p.committee = resp.Peer
	return nil
}

//we process all transactions in block at once, since we need to form proof for all txs in the block^ but not for 2*f+1
func (p *Prover) Prove(ctx context.Context, tpbs []*pb.TransactionS) error {
	addresses := make(map[c.Address]*pb.TransactionS)
	for _, tpb := range tpbs {
		address := c.BytesToAddress(tpb.GetTo())
		list := p.proofs[address]
		list = append(list, tpb.GetData())
		addresses[address] = tpb

	}
	for a, tpb := range addresses {
		list := p.proofs[a]
		if len(list) >= 2*(len(p.committee)-1)/3 {
			mappedProofs := make(map[c.Address]*crypto_g.Signature)
			var signs []*crypto_g.Signature
			for _, proof := range list {
				pbSign := &pb.Signature{}
				if err := proto.Unmarshal(proof, pbSign); err != nil {
					return err
				}
				signFromProto := crypto_g.SignatureFromProto(pbSign)
				signs = append(signs, signFromProto)
				address := crypto_g.PubkeyToAddress(crypto_g.NewPublicKey(signFromProto.Pub()))
				mappedProofs[address] = signFromProto
			}

			bitmap := getBitmap(mappedProofs, p.committee)
			spew.Dump(mappedProofs)
			aggregate := crypto_g.AggregateSignatures(bitmap, signs)
			prAggr := aggregate.ToProto()
			aggrBytes, _ := proto.Marshal(prAggr)

			opts := &tx.TransactOpts{
				From:       c.Address{},
				PrivateKey: p.pk,
				Fee:        big.NewInt(1),
				Context:    ctx,
			}
			//we use tpb value since all values are the same for distinct address
			if err := p.txSend.Transact(opts, api.Proof, a, big.NewInt(tpb.Value), aggrBytes); err != nil {
				return err
			}
			delete(p.proofs, a)
		}
	}

	return nil
}

func getBitmap(src map[c.Address]*crypto_g.Signature, committee []*pb.Peer) *big.Int {
	var addresses []c.Address
	for _, p := range committee {
		addresses = append(addresses, c.BytesToAddress(p.Address))
	}
	return c2.GetBitmap(src, addresses)
}
