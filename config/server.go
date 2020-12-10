package config

import "fmt"

// ServerListen for http & grpc hostname
type ServerListen struct {
	Host string `mapstructure:"host"`
	Port uint16 `mapstructure:"port"`
}

func (s ServerListen) String() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// Server for server configure
type Server struct {
	GRPC ServerListen `mapstructure:"grpc"`
	HTTP ServerListen `mapstructure:"http"`
}
