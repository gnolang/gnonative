module github.com/gnolang/gnonative

go 1.21

toolchain go1.21.7

require (
	connectrpc.com/connect v1.13.0
	connectrpc.com/grpchealth v1.2.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/gnolang/gno v0.0.0-20240306183113-21c754ac4736
	github.com/oklog/run v1.1.0
	github.com/peterbourgon/ff/v3 v3.4.0
	github.com/peterbourgon/unixtransport v0.0.3
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.10.1
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.26.0
	golang.org/x/mobile v0.0.0-20230531173138-3c911d8e3eda
	golang.org/x/net v0.21.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/protobuf v1.31.0
	moul.io/u v1.27.0
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.2 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.3 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cosmos/ledger-cosmos-go v0.13.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/gnolang/goleveldb v0.0.9 // indirect
	github.com/gnolang/overflow v0.0.0-20170615021017-4d914c927216 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/jaekwon/testify v1.6.1 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/linxGnu/grocksdb v1.8.11 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/zondax/hid v0.9.2 // indirect
	github.com/zondax/ledger-go v0.14.3 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231009173412-8bfb1ae86b6c // indirect
	google.golang.org/grpc v1.58.3 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace golang.org/x/mobile => github.com/berty/mobile v0.0.11 // temporary, see https://github.com/golang/mobile/pull/58 and https://github.com/golang/mobile/pull/82
