TRUSTWALLETCORE_IMAGE=trustwalletcore
TRUSTWALLETCORE_DIR=deps/trustwalletcore
PROJECT_ROOT := $(PWD)

CGO_CFLAGS := -I$(PROJECT_ROOT)/deps/trustwalletcore/include

CGO_LDFLAGS := \
-L$(PROJECT_ROOT)/deps/trustwalletcore/build \
-L$(PROJECT_ROOT)/deps/trustwalletcore/build/local/lib \
-L$(PROJECT_ROOT)/deps/trustwalletcore/build/trezor-crypto \
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


.PHONY: install
install:
	docker build \
		-t $(TRUSTWALLETCORE_IMAGE) \
		https://github.com/trustwallet/wallet-core.git

	rm -rf $(TRUSTWALLETCORE_DIR)

	docker create --name tw $(TRUSTWALLETCORE_IMAGE) 
	mkdir -p deps 
	docker cp tw:/wallet-core $(TRUSTWALLETCORE_DIR) 
	docker rm tw

	rm -rf internal/proto
	mkdir -p internal/proto

	cp -R deps/trustwalletcore/samples/go/protos/common internal/proto/
	cp -R deps/trustwalletcore/samples/go/protos/ethereum internal/proto/

	sed -i \
		's|tw/protos/common|wallet-service/internal/proto/common|g' \
		internal/proto/ethereum/Ethereum.pb.go

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