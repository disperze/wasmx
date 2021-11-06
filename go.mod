module github.com/disperze/wasmx

go 1.16

require (
	github.com/CosmWasm/wasmd v0.16.0
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/forbole/juno/v2 v2.0.0-20211020184842-e358a33007ff
	github.com/jmoiron/sqlx v1.3.4
	github.com/proullon/ramsql v0.0.0-20210730175921-2692f3496a21
	github.com/stretchr/testify v1.7.0
)

replace (
	github.com/CosmWasm/wasmd => github.com/notional-labs/wasmd v1.0.0-juno-beta
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)
