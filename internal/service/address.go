package service

import (
	"wallet-service/internal/gate"
	"wallet-service/internal/trustwalletcore"
)

type AddressService struct {
	gates *gate.Manager
}

func NewAddressService(
	gates *gate.Manager,
) *AddressService {
	return &AddressService{
		gates: gates,
	}
}

func (s *AddressService) CreateAddress(
	gateName string,
	account uint32,
	change uint32,
	addressIndex uint32,
) (string, error) {

	mnemonic, err := s.gates.Mnemonic(gateName)
	if err != nil {
		return "", err
	}

	return trustwalletcore.CreateAddress(
		trustwalletcore.CreateAddressRequest{
			Gate:         gateName,
			Mnemonic:     mnemonic,
			Account:      account,
			Change:       change,
			AddressIndex: addressIndex,
		},
	)
}

func (s *AddressService) ValidateAddress(
	gateName string,
	address string,
) (bool, error) {

	return trustwalletcore.ValidateAddress(
		trustwalletcore.ValidateAddressRequest{
			Gate:    gateName,
			Address: address,
		},
	)
}