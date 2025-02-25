package ethComms

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/BitDSM/BitDSM-Node/PodManager"
	"github.com/BitDSM/BitDSM-Node/address"
	"github.com/BitDSM/BitDSM-Node/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/viper"
)

func SubscribeToDepositRequests() {
	// Create a new instance of the contract binding

	oprEthAccount := LoadEthAccount()
	client, err := rpc.Dial(viper.GetString("eth_ws_host"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(client)
	defer ethClient.Close()
	defer client.Close()

	podManagerAddr := common.HexToAddress(viper.GetString("pod_manager_address"))
	podManager, err := PodManager.NewPodManager(podManagerAddr, ethClient)
	if err != nil {
		fmt.Println("Failed to instantiate contract:", err)
		panic(err)
	}

	// Create a channel for the events
	ch := make(chan *PodManager.PodManagerVerifyBitcoinDepositRequest)

	// Create a subscription
	sub, err := podManager.WatchVerifyBitcoinDepositRequest(
		&bind.WatchOpts{Context: context.Background()},
		ch,
		[]common.Address{},
		[]common.Address{oprEthAccount.Address},
	)
	if err != nil {
		fmt.Println("Failed to subscribe to events:", err)
		panic(err)
	}

	fmt.Println("Successfully subscribed to VerifyBitcoinDepositRequest events")

	// Handle events in a loop
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Subscription error deposit:", err)
			time.Sleep(1 * time.Minute)

		case event := <-ch:
			if event.Operator == oprEthAccount.Address {
				handleDepositRequest(event)
			}
		}
	}
}

func SubscribeToWithdrawRequests() {
	// Create a new instance of the contract binding

	oprEthAccount := LoadEthAccount()
	client, err := rpc.Dial(viper.GetString("eth_ws_host"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(client)
	defer ethClient.Close()
	defer client.Close()

	podManagerAddr := common.HexToAddress(viper.GetString("pod_manager_address"))
	podManager, err := PodManager.NewPodManager(podManagerAddr, ethClient)
	if err != nil {
		fmt.Println("Failed to instantiate contract:", err)
		panic(err)
	}

	// Create a channel for the events
	ch := make(chan *PodManager.PodManagerBitcoinWithdrawalPSBTRequest)

	// Create a subscription
	sub, err := podManager.WatchBitcoinWithdrawalPSBTRequest(
		&bind.WatchOpts{Context: context.Background()},
		ch,
		[]common.Address{},
		[]common.Address{oprEthAccount.Address},
	)
	if err != nil {
		fmt.Println("Failed to subscribe to events:", err)
		panic(err)
	}

	fmt.Println("Successfully subscribed to Withdrawal events")

	// Handle events in a loop
	for {
		select {
		case err := <-sub.Err():
			fmt.Println("Subscription error withdrawal:", err)
			time.Sleep(1 * time.Minute)

		case event := <-ch:
			if event.Operator == oprEthAccount.Address {
				HandleWithdrawalRequest(event)
				continue
			}
		}
	}
}

func handleDepositRequest(event *PodManager.PodManagerVerifyBitcoinDepositRequest) {
	fmt.Println("got a deposit event with data : ", event.BitcoinDepositRequest)
	fmt.Println(event.Operator)
	fmt.Println(event.Pod)
	dbconn := db.InitDB()
	defer dbconn.Close()

	if event.BitcoinDepositRequest.IsPending == false {
		return
	}

	transactionId := hex.EncodeToString(event.BitcoinDepositRequest.TransactionId[:])
	db.InsertDepositRequest(dbconn, event.Pod.Hex(), event.Operator.Hex(), transactionId, event.BitcoinDepositRequest.Amount)

}

func HandleWithdrawalRequest(event *PodManager.PodManagerBitcoinWithdrawalPSBTRequest) {
	fmt.Println("got an withdrawal event with data : ", event.WithdrawAddress)
	fmt.Println(event.Operator)
	fmt.Println(event.Pod)
	fmt.Println(event.WithdrawAddress)
	psbt, amount, err := address.GenerateMultisigwithdrawTx(event.WithdrawAddress, event.Pod.Hex())
	if err != nil {
		fmt.Println("Error generating psbt : ", err)
		return
	}

	txid, psbt, err := address.SignMultisigPSBT(psbt)
	if err != nil {
		fmt.Println("Error signing psbt : ", err)
		return
	}

	withdrawAddress := event.WithdrawAddress
	_, err = CallWithdrawBitcoinPSBT(event.Pod.Hex(), withdrawAddress, psbt, *big.NewInt(amount))
	if err != nil {
		fmt.Println("Error in calling withdraw bitcoin psbt : ", err)
		return
	}

	dbconn := db.InitDB()
	defer dbconn.Close()
	addrBytes, _ := hex.DecodeString(event.WithdrawAddress)
	if err != nil {
		fmt.Println("Error converting withdraw address to bytes : ", err)
		return
	}
	db.InsertWithDrawRequest(dbconn, event.Pod.Hex(), event.Operator.Hex(), txid, addrBytes)

}
