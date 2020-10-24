package plugin

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Balance(s *Settings, a string) {
	_, instance, _, _, _ := InitEth(s)
	log.Info(a)
	address := common.HexToAddress(a)
	balance, err := instance.GetBalance(&bind.CallOpts{}, address)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(balance)
}
