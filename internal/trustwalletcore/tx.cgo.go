package trustwalletcore

/*
#include <stdlib.h>

#include <TrustWalletCore/TWHDWallet.h>
#include <TrustWalletCore/TWPrivateKey.h>
#include <TrustWalletCore/TWAnySigner.h>
#include <TrustWalletCore/TWData.h>
#include <TrustWalletCore/TWString.h>
*/
import "C"

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"unsafe"

	ethereum "tw/protos/ethereum"

	"google.golang.org/protobuf/proto"
)

func signEthereumTx(
	req SignTxRequest,
	mnemonic string,
) (string, string, error) {

	privateKeyBytes, err := deriveEthereumPrivateKey(req, mnemonic)
	if err != nil {
		return "", "", err
	}

	input := &ethereum.SigningInput{
		ChainId:   uint64ToBytes(req.ChainID),
		Nonce:     uint64ToBytes(req.Nonce),
		TxMode:    ethereum.TransactionMode_Enveloped,
		GasLimit:  uint64ToBytes(req.GasLimit),
		MaxFeePerGas:          decimalToBytes(req.MaxFeePerGasWei),
		MaxInclusionFeePerGas: decimalToBytes(req.MaxPriorityFeePerGasWei),
		ToAddress: req.To,
		PrivateKey: privateKeyBytes,
		Transaction: &ethereum.Transaction{
			TransactionOneof: &ethereum.Transaction_Transfer_{
				Transfer: &ethereum.Transaction_Transfer{
					Amount: decimalToBytes(req.ValueWei),
					Data:   decodeHex(req.Data),
				},
			},
		},
	}

	inputBytes, err := proto.Marshal(input)
	if err != nil {
		return "", "", err
	}

	twInput := C.TWDataCreateWithBytes(
		(*C.uint8_t)(unsafe.Pointer(&inputBytes[0])),
		C.size_t(len(inputBytes)),
	)
	defer C.TWDataDelete(twInput)

	outputData := C.TWAnySignerSign(
		twInput,
		C.TWCoinTypeEthereum,
	)
	defer C.TWDataDelete(outputData)

	outputBytes := C.GoBytes(
		unsafe.Pointer(C.TWDataBytes(outputData)),
		C.int(C.TWDataSize(outputData)),
	)

	var output ethereum.SigningOutput
	if err := proto.Unmarshal(outputBytes, &output); err != nil {
		return "", "", err
	}

	if output.Error != 0 {
		return "", "", fmt.Errorf(output.ErrorMessage)
	}

	txHash := "0x" + hex.EncodeToString(output.PreHash)
	signedTx := "0x" + hex.EncodeToString(output.Encoded)

	return txHash, signedTx, nil
}

func deriveEthereumPrivateKey(
	req SignTxRequest,
	mnemonic string,
) ([]byte, error) {

	cMnemonic := C.CString(mnemonic)
	defer C.free(unsafe.Pointer(cMnemonic))

	cPassphrase := C.CString("")
	defer C.free(unsafe.Pointer(cPassphrase))

	mn := C.TWStringCreateWithUTF8Bytes(cMnemonic)
	ps := C.TWStringCreateWithUTF8Bytes(cPassphrase)

	defer C.TWStringDelete(mn)
	defer C.TWStringDelete(ps)

	wallet := C.TWHDWalletCreateWithMnemonic(mn, ps)
	if wallet == nil {
		return nil, fmt.Errorf("failed wallet")
	}
	defer C.TWHDWalletDelete(wallet)

	priv := C.TWHDWalletGetDerivedKey(
		wallet,
		C.TWCoinTypeEthereum,
		C.uint32_t(req.Account),
		C.uint32_t(req.Change),
		C.uint32_t(req.AddressIndex),
	)
	if priv == nil {
		return nil, fmt.Errorf("failed key")
	}
	defer C.TWPrivateKeyDelete(priv)

	data := C.TWPrivateKeyData(priv)
	defer C.TWDataDelete(data)

	return C.GoBytes(
		unsafe.Pointer(C.TWDataBytes(data)),
		C.int(C.TWDataSize(data)),
	), nil
}

func uint64ToBytes(v uint64) []byte {
	if v == 0 {
		return []byte{}
	}
	return new(big.Int).SetUint64(v).Bytes()
}

func decimalToBytes(v string) []byte {
	if v == "" {
		return []byte{}
	}
	n := new(big.Int)
	n.SetString(v, 10)
	return n.Bytes()
}

func decodeHex(v string) []byte {
	v = strings.TrimPrefix(v, "0x")
	if v == "" {
		return nil
	}
	b, _ := hex.DecodeString(v)
	return b
}