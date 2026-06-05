package trustwalletcore

type SignTxRequest struct {
	Gate         string
	Account      uint32
	Change       uint32
	AddressIndex uint32
	To                      string
	ValueWei                string
	Data                    string
	Nonce                   uint64
	ChainID                 uint64
	GasLimit                uint64
	MaxFeePerGasWei         string
	MaxPriorityFeePerGasWei string
}

type SignTxResponse struct {
	TxHash   string
	SignedTx string
}

func SignEthereumTx(req SignTxRequest, mnemonic string) (string, string, error) {
	return signEthereumTx(req, mnemonic)
}