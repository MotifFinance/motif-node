package btcComms

type JSONRPCRequest struct {
	ID      int64       `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Jsonrpc string      `json:"jsonrpc"`
}

// generic string response
type JSONRPCResponse struct {
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

type JSONRPCArrayResponse struct {
	Result []string    `json:"result"`
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

type JSONRPCResponseAddressInfo struct {
	Result AddressInfo `json:"result"`
	Error  interface{} `json:"error"`
	ID     int64       `json:"id"`
}

type AddressInfo struct {
	Address        string   `json:"address"`
	ScriptPubKey   string   `json:"scriptPubKey"`
	Ismine         bool     `json:"ismine"`
	Solvable       bool     `json:"solvable"`
	Desc           string   `json:"desc"`
	ParentDesc     string   `json:"parent_desc"`
	IsWatchOnly    bool     `json:"iswatchonly"`
	Isscript       bool     `json:"isscript"`
	IsWitness      bool     `json:"iswitness"`
	WitnessVersion int      `json:"witness_version"`
	WitnessProgram string   `json:"witness_program"`
	Script         string   `json:"script"`
	Hex            string   `json:"hex"`
	IsChange       bool     `json:"ischange"`
	Labels         []string `json:"labels"`
}

// response with desc info
type JSONRPCResponseDesc struct {
	Result DescriptorInfo `json:"result"`
	Error  interface{}    `json:"error"`
	ID     int64          `json:"id"`
}

type DescriptorInfo struct {
	Descriptor     string `json:"descriptor"`
	Checksum       string `json:"checksum"`
	IsRange        bool   `json:"isrange"`
	IsSolvable     bool   `json:"issolvable"`
	HasPrivateKeys bool   `json:"hasprivatekeys"`
}

// //response for importing descriptor
type ImportDescriptorType struct {
	Desc      string `json:"desc"`
	Active    bool   `json:"active"`
	Internal  bool   `json:"internal"`
	Timestamp string `json:"timestamp"`
}

// response for decoding psbt
type JSONRPCResponsePsbt struct {
	Result PSBT        `json:"result"`
	Error  interface{} `json:"error"`
	ID     int64       `json:"id"`
}

type PSBT struct {
	Tx          Tx                `json:"tx"`
	GlobalXPubs []interface{}     `json:"global_xpubs"`
	PSBTVersion int               `json:"psbt_version"`
	Proprietary []interface{}     `json:"proprietary"`
	Unknown     map[string]string `json:"unknown"`
	Inputs      []Input           `json:"inputs"`
	Outputs     []Output          `json:"outputs"`
	Fee         float64           `json:"fee"`
}

type Tx struct {
	TxID     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	VSize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
}

type Vin struct {
	TxID      string    `json:"txid"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  int       `json:"sequence"`
}

type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Desc    string `json:"desc"`
	Hex     string `json:"hex"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

type Input struct {
	WitnessUTXO    WitnessUTXO           `json:"witness_utxo"`
	NonWitnessUTXO NonWitnessUTXO        `json:"non_witness_utxo"`
	PartialSigs    map[string]string     `json:"partial_signatures"`
	WitnessScript  Script                `json:"witness_script"`
	BIP32Derivs    []BIP32DerivationPath `json:"bip32_derivs"`
}

type WitnessUTXO struct {
	Amount       float64      `json:"amount"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type NonWitnessUTXO struct {
	TxID     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	VSize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []Vin  `json:"vin"`
	Vout     []Vout `json:"vout"`
}

type Script struct {
	Asm  string `json:"asm"`
	Hex  string `json:"hex"`
	Type string `json:"type"`
}

type BIP32DerivationPath struct {
	PubKey            string `json:"pubkey"`
	MasterFingerprint string `json:"master_fingerprint"`
	Path              string `json:"path"`
}

type Output struct {
	WitnessScript Script                `json:"witness_script,omitempty"`
	BIP32Derivs   []BIP32DerivationPath `json:"bip32_derivs,omitempty"`
}

// response for Create PSbt
type RPCResponseCreatePsbt struct {
	Result struct {
		Psbt     string `json:"psbt"`
		Complete bool   `json:"complete"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}

type RPCResponseSignPsbt struct {
	Result struct {
		Psbt      string  `json:"psbt"`
		Fee       float64 `json:"fee"`
		ChangePos int     `json:"changepos"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}

// structs for creating a transaction/psbt
type TxInput struct {
	Txid     string `json:"txid"`
	Vout     int64  `json:"vout"`
	Sequence int64  `json:"sequence"`
}

type TxOutput map[string]float64

type CreateTx struct {
	Inputs  []TxInput  `json:"inputs"`
	Outputs []TxOutput `json:"outputs"`
}

type EstimateFeeResponse struct {
	Result struct {
		Feerate float64  `json:"feerate,omitempty"`
		Errors  []string `json:"errors,omitempty"`
		Blocks  int      `json:"blocks"`
	}
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}

type TransactionInfo struct {
	Result string
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

type TransactionError struct {
	TxID      string `json:"txid"`
	Vout      int    `json:"vout"`
	ScriptSig string `json:"scriptSig"`
	Sequence  int    `json:"sequence"`
	Error     string `json:"error"`
}

type SignedTransactionInfo struct {
	Result struct {
		Hex      string             `json:"hex"`
		Complete bool               `json:"complete"`
		Errors   []TransactionError `json:"errors,omitempty"`
	}
	Error interface{} `json:"error"`
	ID    int         `json:"id"`
}
