package trustwalletcore

import "testing"

func TestCreateAndValidateAddress(t *testing.T) {
	address, err := CreateAddress(
		CreateAddressRequest{
			Gate:         "ethereum",
			Mnemonic:     "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
			Account:      0,
			Change:       0,
			AddressIndex: 15,
		},
	)

	if err != nil {
		t.Fatalf("create address failed: %v", err)
	}

	valid, err := ValidateAddress(
		ValidateAddressRequest{
			Gate:    "ethereum",
			Address: address,
		},
	)

	if err != nil {
		t.Fatalf("validate address failed: %v", err)
	}

	if !valid {
		t.Fatal("generated address should be valid")
	}
}
