module github.com/Tokumicn/grpc_tlp

go 1.23.0

require (
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/go-kratos/kratos/v2 v2.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/net v0.39.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240827150818-7e3bb234dfed
	google.golang.org/grpc v1.66.0
	google.golang.org/protobuf v1.34.2
	lego_lib v0.0.0-00010101000000-000000000000
)

require (
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/google/gops v0.3.28 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240823204242-4ba0660f739c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace lego_lib => ../lego_lib
