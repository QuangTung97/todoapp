package mysql

// Config for configuring MySQL
type Config struct {
	Host         string `mapstructure:"host"`
	Port         uint16 `mapstructure:"port"`
	Database     string `mapstructure:"database"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	Options      string `mapstructure:"options"`
}
