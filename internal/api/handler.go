package api

import "wallet-service/internal/service"

type Handler struct {
	addressService *service.AddressService
	transactionService *service.TransactionService
}

func NewHandler(
	addressService *service.AddressService,
	transactionService *service.TransactionService,
) *Handler {
	return &Handler{
		addressService: addressService,
		transactionService: transactionService,
	}
}