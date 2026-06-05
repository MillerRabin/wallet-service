TRUSTWALLETCORE_IMAGE=trustwalletcore
TRUSTWALLETCORE_DIR=deps/trustwalletcore
PROJECT_ROOT := $(PWD)
DOCKER_IMAGE=wallet-service
DOCKER_CONTAINER=wallet-service

CGO_CFLAGS := -I$(PROJECT_ROOT)/$(TRUSTWALLETCORE_DIR)/include

CGO_LDFLAGS := \
-L$(PROJECT_ROOT)/$(TRUSTWALLETCORE_DIR)/build \
-L$(PROJECT_ROOT)/$(TRUSTWALLETCORE_DIR)/build/local/lib \
-L$(PROJECT_ROOT)/$(TRUSTWALLETCORE_DIR)/build/trezor-crypto \
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

	cp -R $(TRUSTWALLETCORE_DIR)/samples/go/protos/common internal/proto/
	cp -R $(TRUSTWALLETCORE_DIR)/samples/go/protos/ethereum internal/proto/

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

.PHONY: docker-build
docker-build:
	docker build \
		-t $(DOCKER_IMAGE) \
		.

.PHONY: docker-run
docker-run:
	docker run \
		--rm \
		--name $(DOCKER_CONTAINER) \
		-p 8000:8000 \
		-v $(PROJECT_ROOT)/config.json:/app/config.json:ro \
		$(DOCKER_IMAGE)

.PHONY: docker-stop
docker-stop:
	docker stop $(DOCKER_CONTAINER)