module github.com/gnolang/gnonative/v4

go 1.23.0

toolchain go1.23.7

require (
	connectrpc.com/connect v1.16.2
	connectrpc.com/grpchealth v1.3.0
	connectrpc.com/grpcreflect v1.2.0
	github.com/gnolang/gno v0.0.0-20250728014947-2e6b56fe460e
	github.com/oklog/run v1.1.0
	github.com/peterbourgon/ff/v3 v3.4.0
	github.com/peterbourgon/unixtransport v0.0.3
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.11.1
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/mobile v0.0.0-20230531173138-3c911d8e3eda
	golang.org/x/net v0.39.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
	google.golang.org/protobuf v1.36.3
	moul.io/u v1.27.0
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.4 // indirect
	github.com/btcsuite/btcd/btcutil v1.1.6 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cosmos/ledger-cosmos-go v0.14.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.25.1 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/sig-0/insertion-queue v0.0.0-20241004125609-6b3ca841346b // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/zondax/hid v0.9.2 // indirect
	github.com/zondax/ledger-go v0.14.3 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.34.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk v1.34.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.opentelemetry.io/proto/otlp v1.5.0 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/term v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	golang.org/x/tools v0.32.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.69.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace golang.org/x/mobile => github.com/berty/mobile v0.0.11 // temporary, see https://github.com/golang/mobile/pull/58 and https://github.com/golang/mobile/pull/82
