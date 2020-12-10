package mysql

import (
	"fmt"
	"net/url"
	"strings"
)

// Config for configuring MySQL
type Config struct {
	Host         string            `mapstructure:"host"`
	Port         uint16            `mapstructure:"port"`
	Database     string            `mapstructure:"database"`
	Username     string            `mapstructure:"username"`
	Password     string            `mapstructure:"password"`
	MaxOpenConns int               `mapstructure:"max_open_conns"`
	MaxIdleConns int               `mapstructure:"max_idle_conns"`
	Options      map[string]string `mapstructure:"options"`
}

// DefaultConfig default values
var DefaultConfig = Config{
	Host:         "localhost",
	Port:         3306,
	Username:     "username",
	Password:     "password",
	Database:     "sample",
	MaxOpenConns: 20,
	MaxIdleConns: 5,
	Options: map[string]string{
		"parseTime": "true",
		"loc":       "Asia/Ho_Chi_Minh",
	},
}

// DSN returns data source name
func (c Config) DSN() string {
	var opts []string
	for key, value := range c.Options {
		key = url.QueryEscape(key)
		value = url.QueryEscape(value)
		opts = append(opts, key+"="+value)
	}
	optStr := strings.Join(opts, "&")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, optStr)
}
