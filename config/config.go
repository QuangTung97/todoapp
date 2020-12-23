package config

import (
	"fmt"
	"github.com/spf13/viper"
	"todoapp/lib/log"
	"todoapp/lib/mysql"
)

// Config for configuring whole app
type Config struct {
	Server Server       `mapstructure:"server"`
	Event  Event        `mapstructure:"event"`
	Log    log.Config   `mapstructure:"log"`
	MySQL  mysql.Config `mapstructure:"mysql"`
}

// Load config from config.yml
func Load() Config {
	vip := viper.New()

	vip.SetConfigName("config")
	vip.SetConfigType("yml")
	vip.AddConfigPath(".")

	err := vip.ReadInConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println("Config file used:", vip.ConfigFileUsed())

	cfg := Config{}
	err = vip.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
