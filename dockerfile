# ==================================================
# Trust Wallet Core
# ==================================================
FROM trustwalletcore AS walletcore

# ==================================================
# Build
# ==================================================
FROM golang:1.26.4-trixie AS builder

RUN apt-get update && apt-get install -y \
  clang \
  gcc \
  g++ \
  make \
  libc++-dev \
  libc++abi-dev

ENV CGO_ENABLED=1

COPY --from=walletcore /wallet-core/include /opt/wallet-core/include
COPY --from=walletcore /wallet-core/build /opt/wallet-core/build
COPY --from=walletcore /wallet-core/samples/go /opt/wallet-core/samples/go

ENV CGO_CFLAGS="-I/opt/wallet-core/include"

ENV CGO_LDFLAGS="\
  -L/opt/wallet-core/build \
  -L/opt/wallet-core/build/local/lib \
  -L/opt/wallet-core/build/trezor-crypto \
  -lTrustWalletCore \
  -lwallet_core_rs \
  -lTrezorCrypto \
  -lprotobuf \
  -lprotobuf-lite \
  -lstdc++ \
  -lpthread \
  -ldl \
  -lm"

WORKDIR /app

COPY go.mod .
COPY go.sum .


RUN go mod edit -replace=tw=/opt/wallet-core/samples/go
RUN go mod download

COPY . .

RUN go mod edit -replace=tw=/opt/wallet-core/samples/go

RUN go build \
  -ldflags="-s -w" \
  -o /out/wallet-service \
  ./cmd/server

# ==================================================
# Runtime
# ==================================================
FROM debian:trixie-slim

RUN apt-get update && apt-get install -y \
  libstdc++6 \
  && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder \
  /out/wallet-service \
  /wallet-service

ENTRYPOINT ["/wallet-service"]