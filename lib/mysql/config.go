package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/url"
	"strings"
)

// OptionConfig for options
type OptionConfig struct {
	Key   string `mapstructure:"key"`
	Value string `mapstructure:"value"`
}

// Config for configuring MySQL
type Config struct {
	Host         string         `mapstructure:"host"`
	Port         uint16         `mapstructure:"port"`
	Database     string         `mapstructure:"database"`
	Username     string         `mapstructure:"username"`
	Password     string         `mapstructure:"password"`
	MaxOpenConns int            `mapstructure:"max_open_conns"`
	MaxIdleConns int            `mapstructure:"max_idle_conns"`
	Options      []OptionConfig `mapstructure:"options"`
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
	Options: []OptionConfig{
		{Key: "parseTime", Value: "true"},
		{Key: "loc", Value: "Asia/Ho_Chi_Minh"},
	},
}

// DSN returns data source name
func (c Config) DSN() string {
	var opts []string
	for _, o := range c.Options {
		key := url.QueryEscape(o.Key)
		value := url.QueryEscape(o.Value)
		opts = append(opts, key+"="+value)
	}
	optStr := strings.Join(opts, "&")
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Database, optStr)
}

// MustConnect connects to database using sqlx
// MUST import
// _ "github.com/go-sql-driver/mysql"
// to use in main.go
func MustConnect(conf Config) *sqlx.DB {
	db := sqlx.MustConnect("mysql", conf.DSN())

	fmt.Println("MaxOpenConns:", conf.MaxOpenConns)
	fmt.Println("MaxIdleConns:", conf.MaxIdleConns)

	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	return db
}
