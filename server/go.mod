module github.com/kurtosis-tech/kurtosis-package-indexer/server

go 1.19

replace github.com/kurtosis-tech/kurtosis-package-indexer/api/golang => ../api/golang

require (
	connectrpc.com/connect v1.11.0
	github.com/google/go-github/v54 v54.0.0
	github.com/kurtosis-tech/kurtosis-package-indexer/api/golang v0.0.0-00010101000000-000000000000
	github.com/kurtosis-tech/kurtosis/connect-server v0.0.0-20230825003324-75d481e0db8c
	github.com/kurtosis-tech/stacktrace v0.0.0-20211028211901-1c67a77b5409
	github.com/rs/cors v1.9.0
	github.com/sirupsen/logrus v1.9.3
	github.com/stretchr/testify v1.8.4
	go.starlark.net v0.0.0-20230814145427-12f4cb8177e4
	golang.org/x/oauth2 v0.11.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/ProtonMail/go-crypto v0.0.0-20230217124315-7d5c6f04bbb8 // indirect
	github.com/aws/aws-sdk-go v1.44.334 // indirect
	github.com/cloudflare/circl v1.3.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/grpc v1.57.0 // indirect
)
