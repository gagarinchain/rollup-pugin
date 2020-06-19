package plugin

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	pb "github.com/gagarinchain/common/protobuff"
	"github.com/gagarinchain/common/rpc"
	"time"
)

func Send(address string) {
	client := rpc.InitOnNewBlockCreatedClient(address)
	timeout, _ := context.WithTimeout(context.Background(), 10*time.Second)
	created, err := client.Pbc().OnNewBlockCreated(timeout, &pb.OnNewBlockCreatedRequest{
		Block:    nil,
		Receipts: nil,
	})
	if err != nil {
		log.Error(err)
		return
	}
	spew.Dump(created)

}
