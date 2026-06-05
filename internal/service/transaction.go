package service

import "wallet-service/internal/trustwalletcore"

type TransactionService struct {
	gates *GateManager
}

func NewTransactionService(
	gates *GateManager,
) *TransactionService {
	return &TransactionService{
		gates: gates,
	}
}

func (s *TransactionService) SignEthereumTx(
	req trustwalletcore.SignTxRequest,
) (
	string,
	string,
	error,
) {
	mnemonic, err := s.gates.Mnemonic(
		req.Gate,
	)

	if err != nil {
		return "", "", err
	}

	return trustwalletcore.SignEthereumTx(
		req,
		mnemonic,
	)
}