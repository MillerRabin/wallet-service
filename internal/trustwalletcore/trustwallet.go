package trustwalletcore

import "fmt"

const (
	GateEthereum = "ethereum"
)

type CreateAddressRequest struct {
	Gate         string
	Mnemonic     string
	Account      uint32
	Change       uint32
	AddressIndex uint32
}

type ValidateAddressRequest struct {
	Gate    string
	Address string
}


func CreateAddress(
	req CreateAddressRequest,
) (string, error) {
	switch req.Gate {
		case GateEthereum:
			return createEthereumAddress(req)

		default:
			return "", fmt.Errorf(
				"unsupported gate: %s",
				req.Gate,
		)
	}
}

func CreateEthereumAddress(
	req CreateAddressRequest,
) (string, error) {
	req.Gate = GateEthereum
	return createEthereumAddress(req)
}

func ValidateAddress(
	req ValidateAddressRequest,
) (bool, error) {
	return validateAddress(req)
}