package dto

type ValidateAddressRequest struct {
	Gate    string `json:"gate"`
	Address string `json:"address"`
}

type ValidateAddressResponse struct {
	Valid bool `json:"valid"`
}