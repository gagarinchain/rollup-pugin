module github.com/gagarinchain/rollup-plugin

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/ethereum/go-ethereum v1.9.14
	github.com/gagarinchain/common v0.1.10
	github.com/golang/protobuf v1.4.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/opentracing/opentracing-go v1.1.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.0
	google.golang.org/grpc v1.21.1
)

//replace github.com/gagarinchain/common => ../common

go 1.14
