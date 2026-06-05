package dto

type TxRequest struct {
	Gate         string   `json:"gate"`
	Account      uint32   `json:"account"`
	Change       uint32   `json:"change"`
	AddressIndex uint32   `json:"address_index"`
	TxParams     TxParams `json:"tx_params"`
}

type TxParams struct {
	To                       string `json:"to"`
	ValueWei                 string `json:"value_wei"`
	Data                     string `json:"data"`
	Nonce                    uint64 `json:"nonce"`
	ChainID                  uint64 `json:"chain_id"`
	GasLimit                 uint64 `json:"gas_limit"`
	MaxFeePerGasWei          string `json:"max_fee_per_gas_wei"`
	MaxPriorityFeePerGasWei  string `json:"max_priority_fee_per_gas_wei"`
}

type TxResponse struct {
	TxHash   string `json:"tx_hash"`
	SignedTx string `json:"signed_tx"`
}