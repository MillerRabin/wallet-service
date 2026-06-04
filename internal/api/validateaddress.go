package api

import (
	"encoding/json"
	"net/http"

	"wallet-service/internal/dto"
)

func (h *Handler) ValidateAddress(
	w http.ResponseWriter,
	r *http.Request,
) {
	var req dto.ValidateAddressRequest

	if err := json.NewDecoder(r.Body).
		Decode(&req); err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	valid, err := h.addressService.
		ValidateAddress(
			req.Gate,
			req.Address,
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	resp := dto.ValidateAddressResponse{
		Valid: valid,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	_ = json.NewEncoder(w).
		Encode(resp)
}