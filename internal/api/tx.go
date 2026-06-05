package api

import (
	"encoding/json"
	"net/http"

	"wallet-service/internal/service"
	"wallet-service/internal/dto"
	"wallet-service/internal/trustwalletcore"
)

func (h *Handler) SignTx(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.TxRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	// use a pointer receiver for TransactionService
	svc := &service.TransactionService{}

	txHash,
	signedTx,
	err := svc.
		SignEthereumTx(
			trustwalletcore.SignTxRequest{
				Gate:         req.Gate,
				Account:      req.Account,
				Change:       req.Change,
				AddressIndex: req.AddressIndex,

				To:                      req.TxParams.To,
				ValueWei:                req.TxParams.ValueWei,
				Data:                    req.TxParams.Data,
				Nonce:                   req.TxParams.Nonce,
				ChainID:                 req.TxParams.ChainID,
				GasLimit:                req.TxParams.GasLimit,
				MaxFeePerGasWei:         req.TxParams.MaxFeePerGasWei,
				MaxPriorityFeePerGasWei: req.TxParams.MaxPriorityFeePerGasWei,
			},
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	resp := dto.TxResponse{
		TxHash:   txHash,
		SignedTx: signedTx,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).
		Encode(resp)
}