package main

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/MotifFinance/motif-node/BtcDepositConfirmer"
	"github.com/MotifFinance/motif-node/api"
	"github.com/MotifFinance/motif-node/ethComms"
	"github.com/MotifFinance/motif-node/operator"
	"github.com/MotifFinance/motif-node/utils"
)

func initialize() {
	utils.InitConfigFile()
	env := viper.GetString("env")
	if env != "dev" && env != "prod" {
		fmt.Println("Invalid environment")
		panic("Invalid environment")
	}

	utils.LoadBtcWallet(viper.GetString("wallet_name"))
	ethAccount := ethComms.LoadEthAccount()
	fmt.Println("Eth account: ", ethAccount.Address.Hex())
	operator.RegisterOperator()
}

func main() {
	initialize()
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Server()
	wg.Add(1)
	go BtcDepositConfirmer.CheckDeposit()
	wg.Add(1)
	go ethComms.SubscribeToDepositRequests()
	wg.Add(1)
	go BtcDepositConfirmer.CheckWithdraw()
	ethComms.SubscribeToWithdrawRequests()
	wg.Wait()
}
