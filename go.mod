module go-lib

go 1.18

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/shopspring/decimal v1.3.1
	github.com/speps/go-hashids v2.0.0+incompatible
    github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/cosmos/cosmos-sdk v0.39.2
	github.com/ethereum/go-ethereum v1.10.8
	github.com/okex/exchain v0.19.5
	github.com/pkg/errors v0.9.1
	github.com/speps/go-hashids v2.0.0+incompatible
	github.com/tendermint/tendermint v0.33.9
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/okex/cosmos-sdk v0.39.2-exchain16
	github.com/tendermint/iavl => github.com/okex/iavl v0.14.3-exchain1
	github.com/tendermint/tendermint => github.com/okex/tendermint v0.33.9-exchain11
	github.com/tendermint/tm-db => github.com/okex/tm-db v0.5.2-exchain1
)
