package dto

type CreateAddressRequest struct {
	Gate         string `json:"gate"`
	Account      uint32 `json:"account"`
	Change       uint32 `json:"change"`
	AddressIndex uint32 `json:"address_index"`
}

type CreateAddressResponse struct {
	Address string `json:"address"`
}