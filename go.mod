module github.com/gagarinchain/rollup-plugin

require (
	github.com/btcsuite/btcd v0.21.0-beta // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/ethereum/go-ethereum v1.9.20
	github.com/ferranbt/fastssz v0.0.0-20200826142241-3a913c5a1313 // indirect
	github.com/gagarinchain/common v0.1.17
	github.com/gagarinchain/network v0.1.2
	github.com/golang/protobuf v1.4.2
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/minio/highwayhash v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/phoreproject/bls v0.0.0-20200525203911-a88a5ae26844 // indirect
	github.com/protolambda/zssz v0.1.5 // indirect
	github.com/prysmaticlabs/go-bitfield v0.0.0-20200618145306-2ae0807bef65 // indirect
	github.com/prysmaticlabs/go-ssz v0.0.0-20200612203617-6d5c9aa213ae
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sys v0.0.0-20200908134130-d2e65c121b96 // indirect
	google.golang.org/grpc v1.33.0
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/gagarinchain/common => ../common

replace github.com/gagarinchain/network => ../network

go 1.14
