package trustwalletcore

/*
#include <stdlib.h>

#include <TrustWalletCore/TWHDWallet.h>
#include <TrustWalletCore/TWString.h>
#include <TrustWalletCore/TWCoinType.h>
#include <TrustWalletCore/TWPrivateKey.h>
#include <TrustWalletCore/TWPublicKey.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func createEthereumAddress(
	req CreateAddressRequest,
) (string, error) {
	if req.Mnemonic == "" {
		return "", fmt.Errorf("mnemonic is required")
	}

	wallet, err := createWallet(req.Mnemonic)
	if err != nil {
		return "", err
	}
	defer C.TWHDWalletDelete(wallet)

	return deriveAddress(
		wallet,
		C.TWCoinTypeEthereum,
		req.Account,
		req.Change,
		req.AddressIndex,
	)
}

func createWallet(
	mnemonic string,
) (*C.struct_TWHDWallet, error) {
	cMnemonic := C.CString(mnemonic)
	defer C.free(unsafe.Pointer(cMnemonic))

	cPassphrase := C.CString("")
	defer C.free(unsafe.Pointer(cPassphrase))

	twMnemonic := C.TWStringCreateWithUTF8Bytes(cMnemonic)
	defer C.TWStringDelete(twMnemonic)

	twPassphrase := C.TWStringCreateWithUTF8Bytes(cPassphrase)
	defer C.TWStringDelete(twPassphrase)

	wallet := C.TWHDWalletCreateWithMnemonic(
		twMnemonic,
		twPassphrase,
	)

	if wallet == nil {
		return nil, fmt.Errorf("failed to create wallet")
	}

	return wallet, nil
}

func deriveAddress(
	wallet *C.struct_TWHDWallet,
	coinType C.enum_TWCoinType,
	account uint32,
	change uint32,
	addressIndex uint32,
) (string, error) {
	privateKey := C.TWHDWalletGetDerivedKey(
		wallet,
		coinType,
		C.uint32_t(account),
		C.uint32_t(change),
		C.uint32_t(addressIndex),
	)

	if privateKey == nil {
		return "", fmt.Errorf("failed to derive private key")
	}
	defer C.TWPrivateKeyDelete(privateKey)

	publicKey := C.TWPrivateKeyGetPublicKey(
		privateKey,
		coinType,
	)

	if publicKey == nil {
		return "", fmt.Errorf("failed to derive public key")
	}
	defer C.TWPublicKeyDelete(publicKey)

	address := C.TWCoinTypeDeriveAddressFromPublicKey(
		coinType,
		publicKey,
	)

	if address == nil {
		return "", fmt.Errorf("failed to derive address")
	}
	defer C.TWStringDelete(address)

	return C.GoString(
		C.TWStringUTF8Bytes(address),
	), nil
}

func validateAddress(
	req ValidateAddressRequest,
) (bool, error) {

	cAddress := C.CString(
		req.Address,
	)
	defer C.free(
		unsafe.Pointer(cAddress),
	)

	address := C.TWStringCreateWithUTF8Bytes(
		cAddress,
	)
	defer C.TWStringDelete(
		address,
	)

	valid := C.TWCoinTypeValidate(
		C.TWCoinTypeEthereum,
		address,
	)

	return bool(valid), nil
}