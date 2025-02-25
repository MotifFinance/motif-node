package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/BitDSM/BitDSM-Node/types"
	"github.com/spf13/viper"
)

func InitDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", viper.Get("DB_host"), viper.Get("DB_port"), viper.Get("DB_user"), viper.Get("DB_password"), viper.Get("DB_name"))
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Println("DB error : ", err)
		panic(err)
	}
	fmt.Println("DB initialized")
	return db
}

func UpdateMultiSigAddressPod(dbconn *sql.DB, address string, podaddress string) error {
	_, err := dbconn.Exec("UPDATE multi_sig_address SET podaddress = $1 WHERE address = $2",
		podaddress,
		address,
	)
	if err != nil {
		fmt.Println("An error occurred while executing update multi sig address query: ", err)
		return err
	}
	return nil
}

func InsertMultiSigAddress(dbconn *sql.DB, address string, script string, ethAddr string) error {
	_, err := dbconn.Exec("INSERT into multi_sig_address VALUES ($1, $2, $3, $4, $5)",
		address,
		script,
		nil,
		false,
		false,
	)
	if err != nil {
		fmt.Println("An error occured while executing insert multi sig address query: ", err)
		return err
	}
	return nil
}

func QueryMultisigAddressByPodAddress(dbconn *sql.DB, podAddr string) []types.MultiSigAddress {
	// fmt.Println("getting address for height: ", height)
	var DB_reader *sql.Rows
	var err error
	DB_reader, err = dbconn.Query("select * from multi_sig_address where podaddress = $1 and archived = false", podAddr)

	if err != nil {
		fmt.Println("An error occured while query address by Eth Address: ", err)
	}

	defer DB_reader.Close()
	addresses := make([]types.MultiSigAddress, 0)

	for DB_reader.Next() {
		address := types.MultiSigAddress{}
		err := DB_reader.Scan(
			&address.Address,
			&address.Script,
			&address.PodAddress,
			&address.Signed,
			&address.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		addresses = append(addresses, address)
	}

	return addresses
}

func QueryMultisigAddresses(dbconn *sql.DB) []types.MultiSigAddress {
	// fmt.Println("getting address for height: ", height)
	var DB_reader *sql.Rows
	var err error
	var podAddress sql.NullString
	DB_reader, err = dbconn.Query("select * from multi_sig_address where archived = false")

	if err != nil {
		fmt.Println("An error occured while query address: ", err)
	}

	defer DB_reader.Close()
	addresses := make([]types.MultiSigAddress, 0)

	for DB_reader.Next() {
		address := types.MultiSigAddress{}
		err := DB_reader.Scan(
			&address.Address,
			&address.Script,
			&podAddress,
			&address.Signed,
			&address.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		if podAddress.Valid {
			address.PodAddress = podAddress.String
		} else {
			address.PodAddress = ""
		}

		addresses = append(addresses, address)
	}

	return addresses
}

func MarkMultisigProcessed(dbconn *sql.DB, address string) {
	_, err := dbconn.Exec("UPDATE multi_sig_address SET archived = true WHERE address = $1", address)
	if err != nil {
		fmt.Println("An error occured while updating multi sig address query: ", err)
	}
}

func InsertDepositRequest(dbconn *sql.DB, podAddr string, operatorAddr string, txid string, amount *big.Int) {
	amountStr := amount.String()
	_, err := dbconn.Exec("INSERT into deposit_requests VALUES ($1, $2, $3, $4, $5)",
		podAddr,
		operatorAddr,
		txid,
		amountStr,
		false,
	)
	if err != nil {
		fmt.Println("An error occured while insert deposit request query: ", err)
	}
}

func QueryDepositRequests(dbconn *sql.DB) []types.BtcDepositRequest {
	DB_reader, err := dbconn.Query("select * from deposit_requests where Archived = false")
	if err != nil {
		fmt.Println("An error occured while query deposit request: ", err)
		return nil
	}

	defer DB_reader.Close()
	depositRequests := make([]types.BtcDepositRequest, 0)

	for DB_reader.Next() {
		depositRequest := types.BtcDepositRequest{}
		err := DB_reader.Scan(
			&depositRequest.PodAddress,
			&depositRequest.OperatorAddress,
			&depositRequest.TransactionID,
			&depositRequest.Amount,
			&depositRequest.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		depositRequests = append(depositRequests, depositRequest)
	}
	return depositRequests
}

func MarkDepositRequestAsConfirmed(dbconn *sql.DB, txid string) {
	_, err := dbconn.Exec("UPDATE deposit_requests SET archived = true WHERE transaction_id = $1", txid)
	if err != nil {
		fmt.Println("An error occured while updating deposit request query: ", err)
	}
}

func InsertWithDrawRequest(dbconn *sql.DB, podAddr string, operatorAddr string, txid string, withdrawAddr []byte) {
	_, err := dbconn.Exec("INSERT into withdraw_requests VALUES ($1, $2, $3, $4, $5)",
		podAddr,
		operatorAddr,
		txid,
		withdrawAddr,
		false,
	)
	if err != nil {
		fmt.Println("An error occured while insert withdraw request query: ", err)
	}
}

func MarkWithdrawRequestAsConfirmed(dbconn *sql.DB, txid string) {
	_, err := dbconn.Exec("UPDATE withdraw_requests SET archived = true WHERE transactionid = $1", txid)
	if err != nil {
		fmt.Println("An error occured while updating deposit request query: ", err)
	}
}

func QueryWithdrawRequests(dbconn *sql.DB) []types.BtcWithDrawRequest {
	DB_reader, err := dbconn.Query("select * from withdraw_requests where Archived = false")
	if err != nil {
		fmt.Println("An error occured while query withdraw request: ", err)
		return nil
	}

	defer DB_reader.Close()
	withdrawRequests := make([]types.BtcWithDrawRequest, 0)

	for DB_reader.Next() {
		withdrawRequest := types.BtcWithDrawRequest{}
		err := DB_reader.Scan(
			&withdrawRequest.PodAddress,
			&withdrawRequest.OperatorAddress,
			&withdrawRequest.TransactionID,
			&withdrawRequest.WithdrawAddr,
			&withdrawRequest.Archived,
		)
		if err != nil {
			fmt.Println(err)
		}
		withdrawRequests = append(withdrawRequests, withdrawRequest)
	}
	return withdrawRequests
}
