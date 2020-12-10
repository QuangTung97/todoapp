module todoapp

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1
	github.com/golang-migrate/migrate/v4 v4.14.1
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/kisielk/errcheck v1.2.0
	github.com/prometheus/client_golang v0.9.3
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b
	google.golang.org/grpc v1.34.0
	gopkg.in/yaml.v2 v2.4.0
	todoapp-rpc v0.0.0
)

replace todoapp-rpc v0.0.0 => ../todoapp-rpc/
