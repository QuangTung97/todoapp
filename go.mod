module todoapp

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kisielk/errcheck v1.2.0
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/prometheus/client_golang v0.9.3
	github.com/sahilm/fuzzy v0.1.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/sys v0.0.0-20201207223542-d4d67f95c62d // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/genproto v0.0.0-20201209185603-f92720507ed4 // indirect
	google.golang.org/grpc v1.34.0
	gopkg.in/yaml.v2 v2.4.0
	todoapp-rpc v0.0.0
)

replace todoapp-rpc v0.0.0 => ../todoapp-rpc/
