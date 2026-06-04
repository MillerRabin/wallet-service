package api

import "wallet-service/internal/service"

type Handler struct {
	addressService *service.AddressService
}

func NewHandler(
	addressService *service.AddressService,
) *Handler {
	return &Handler{
		addressService: addressService,
	}
}