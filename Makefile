PROJECT_ROOT := $(PWD)

CGO_CFLAGS := -I$(PROJECT_ROOT)/.trustwallet-core/include

CGO_LDFLAGS := \
-L$(PROJECT_ROOT)/.trustwallet-core/build \
-L$(PROJECT_ROOT)/.trustwallet-core/build/local/lib \
-L$(PROJECT_ROOT)/.trustwallet-core/build/trezor-crypto \
-lTrustWalletCore \
-lwallet_core_rs \
-lTrezorCrypto \
-lprotobuf \
-lprotobuf-lite \
-lstdc++ \
-lpthread \
-ldl \
-lm

export CGO_CFLAGS
export CGO_LDFLAGS

.PHONY: build
build:
	go build \
	-o bin/wallet-service \
	./cmd/server

.PHONY: run
run:
	go run ./cmd/server

.PHONY: test
test:
	go test ./... -v

.PHONY: test-race
test-race:
	go test ./... -race -v

.PHONY: clean
clean:
	rm -rf bin

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: tidy
tidy:
	go mod tidy