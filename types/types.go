package types

type WatchtowerTxInput struct {
	Address string
	Amount  uint64
	Txid    string
	Vout    uint32
}

type MultiSigAddress struct {
	Address    string
	Script     string
	PodAddress string
	Signed     bool
	Archived   bool
}

type Utxo struct {
	Txid   string
	Vout   uint32
	Amount uint64
}

type BtcPubkeyArgs struct {
	BTCPubKey string
	EthAddr   string
}

type BtcDepositRequest struct {
	PodAddress      string // Ethereum address of the pod
	OperatorAddress string // Ethereum address of the operator
	TransactionID   string // Transaction ID of the Bitcoin deposit
	Amount          string // Amount of the deposit
	Archived        bool   // Status of the deposit request
}

type BtcWithDrawRequest struct {
	PodAddress      string // Ethereum address of the pod
	OperatorAddress string // Ethereum address of the operator
	TransactionID   string // Transaction ID of the Bitcoin deposit
	WithdrawAddr    []byte // Amount of the deposit
	Archived        bool   // Status of the deposit request
}
