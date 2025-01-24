module github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/skywalkingencodingextension

go 1.22.0

require (
	github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/skywalking v0.117.0
	github.com/stretchr/testify v1.10.0
	go.opentelemetry.io/collector/component v0.117.1-0.20250114172347-71aae791d7f8
	go.opentelemetry.io/collector/component/componenttest v0.117.0
	go.opentelemetry.io/collector/extension v0.117.1-0.20250114172347-71aae791d7f8
	go.opentelemetry.io/collector/extension/extensiontest v0.117.0
	go.opentelemetry.io/collector/pdata v1.23.1-0.20250114172347-71aae791d7f8
	google.golang.org/protobuf v1.36.2
	skywalking.apache.org/repo/goapi v0.0.0-20240104145220-ba7202308dd4
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.117.1-0.20250114172347-71aae791d7f8 // indirect
	go.opentelemetry.io/collector/semconv v0.117.1-0.20250114172347-71aae791d7f8 // indirect
	go.opentelemetry.io/otel v1.32.0 // indirect
	go.opentelemetry.io/otel/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk v1.32.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.32.0 // indirect
	go.opentelemetry.io/otel/trace v1.32.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/skywalking => ../../../pkg/translator/skywalking