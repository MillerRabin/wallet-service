package trustwalletcore

import (
	"testing"
)

func TestSignEthereumTx(t *testing.T) {
	req := SignTxRequest{
		Gate:         "ethereum",
		Account:      0,
		Change:       0,
		AddressIndex: 0,
	}

	var mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"

	req.To = "0x000000000000000000000000000000000000dead"
	req.ValueWei = "1000000000000000" // 0.001 ETH
	req.Data = ""
	req.Nonce = 1
	req.ChainID = 11155111 // Sepolia
	req.GasLimit = 21000
	req.MaxFeePerGasWei = "30000000000"
	req.MaxPriorityFeePerGasWei = "1500000000"

	txHash, signedTx, err := SignEthereumTx(req, mnemonic)
	if err != nil {
		t.Fatalf("SignEthereumTx() error = %v", err)
	}

	if txHash == "" {
		t.Fatal("txHash is empty")
	}

	if signedTx == "" {
		t.Fatal("signedTx is empty")
	}

	if signedTx[:2] != "0x" {
		t.Fatalf("unexpected signed tx: %s", signedTx)
	}
}