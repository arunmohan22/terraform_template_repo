module {{cookiecutter.__repository_url}}

go 1.18

require (
	github.com/google/wire v0.5.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	github.hpe.com/cloud/go-gadgets/x/config v0.0.0-20220929113716-9b6208496458
	github.hpe.com/cloud/go-gadgets/x/logging v0.0.3
	{%- if cookiecutter.grpc == "True" %}
	github.hpe.com/cloud/go-gadgets/x/grpcserver v0.0.3
	github.hpe.com/cloud/storage-proto v0.0.429
	{%- endif %}
	{%- if cookiecutter.rest == "True" %}
	github.com/go-chi/chi/v5 v5.0.7
	github.hpe.com/cloud/go-gadgets/x/headers v0.0.3
	github.hpe.com/cloud/go-gadgets/x/restserver v0.0.1
	github.hpe.com/cloud/go-gadgets/x/restserver/chiserver v0.0.1
	github.hpe.com/cloud/go-gadgets/x/tracing v0.0.3
	go.opentelemetry.io/otel v1.11.0
	{%- endif %}
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.hpe.com/cloud/go-gadgets/x/jwt v0.0.2 // indirect
	github.hpe.com/cloud/go-gadgets/x/tock v0.0.2 // indirect
	go.opentelemetry.io/otel/sdk v1.9.0 // indirect
	go.opentelemetry.io/otel/trace v1.9.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.5.1 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/knadh/koanf v1.4.3 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.hpe.com/cloud/go-gadgets/x/headers v0.0.3 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.hpe.com/cloud/go-gadgets/x/terrors v0.0.2 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37 // indirect
	google.golang.org/grpc v1.52.0
	google.golang.org/protobuf v1.28.1 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	{%- if cookiecutter.rest == "True" %}
	github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220314183648-97c793e446ba // indirect
	github.com/cocoonspace/dynjson v1.0.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.hpe.com/cloud/go-gadgets/x/jwt v0.0.2 // indirect
	github.hpe.com/cloud/go-gadgets/x/tock v0.0.2 // indirect
	github.hpe.com/cloud/odata-filter-go v0.2.0 // indirect
	go.opentelemetry.io/otel/sdk v1.11.0 // indirect
	go.opentelemetry.io/otel/trace v1.11.0 // indirect
	{%- endif %}
)
