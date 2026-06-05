package api

import (
	"encoding/json"
	"net/http"
	"wallet-service/internal/dto"
)

func (h *Handler) CreateAddress(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.CreateAddressRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	address, err := h.addressService.
		CreateAddress(
			req.Gate,
			req.Account,
			req.Change,
			req.AddressIndex,
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	resp := dto.CreateAddressResponse{
		Address: address,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).
		Encode(resp)
}