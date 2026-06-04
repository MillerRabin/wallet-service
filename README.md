# build
docker build -t trustwalletcore https://github.com/trustwallet/wallet-core.git
docker build .

# local debug

docker create --name tw trustwalletcore
docker cp tw:/wallet-core ./.trustwallet-core
docker rm tw
