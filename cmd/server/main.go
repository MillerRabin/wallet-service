package main

import (
	"fmt"
	"log"
	"net/http"

	"wallet-service/internal/api"
	"wallet-service/internal/config"
	"wallet-service/internal/gate"
	"wallet-service/internal/service"
)

func main() {

	cfg, err := config.Load(
		"config.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	gates := gate.New(
		cfg.Gates,
	)

	addressService := service.NewAddressService(
		gates,
	)

	handler := api.NewHandler(
		addressService,
	)

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /api/v1/createaddress",
		handler.CreateAddress,
	)

	mux.HandleFunc(
	"POST /api/v1/validateaddress",
	handler.ValidateAddress,
)

	addr := fmt.Sprintf(
		"%s:%d",
		cfg.Config.Host,
		cfg.Config.Port,
	)

	fmt.Println(
		"wallet-service listening on",
		addr,
	)

	log.Fatal(
		http.ListenAndServe(
			addr,
			mux,
		),
	)
}