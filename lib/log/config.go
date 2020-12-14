package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config for log configure
type Config struct {
	Level           string   `mapstructure:"level"`
	Mode            string   `mapstructure:"mode"`
	Encoding        string   `mapstructure:"encoding"`
	StacktraceLevel string   `mapstructure:"stacktrace_level"`
	MaskedFields    []string `mapstructure:"masked_fields"`
}

var levelMap = map[string]zapcore.Level{
	"debug":  zap.DebugLevel,
	"info":   zap.InfoLevel,
	"warn":   zap.WarnLevel,
	"error":  zap.ErrorLevel,
	"dpanic": zap.DPanicLevel,
	"panic":  zap.PanicLevel,
	"fatal":  zap.FatalLevel,
}

const (
	production  = "production"
	development = "development"

	json    = "json"
	console = "console"
)

func validateConfig(conf Config) {
	_, ok := levelMap[conf.Level]
	if !ok {
		panic("Invalid log level")
	}

	if conf.Mode != production && conf.Mode != development {
		panic("Invalid log mode")
	}

	if conf.Encoding != json && conf.Encoding != console {
		panic("Invalid log encoding")
	}

	if conf.StacktraceLevel != "" {
		_, ok := levelMap[conf.StacktraceLevel]
		if !ok {
			panic("Invalid log stacktrace_level")
		}
	}
}

// NewLogger creates a zap logger
func NewLogger(conf Config) *zap.Logger {
	validateConfig(conf)

	level := zap.NewAtomicLevelAt(levelMap[conf.Level])

	zapCfg := zap.NewProductionConfig()
	if conf.Mode == development {
		zapCfg = zap.NewDevelopmentConfig()
	}

	zapCfg.Encoding = conf.Encoding
	zapCfg.Level = level

	var option []zap.Option
	if conf.StacktraceLevel != "" {
		stackLevel := zap.NewAtomicLevelAt(levelMap[conf.StacktraceLevel])
		option = append(option, zap.AddStacktrace(stackLevel))
	}

	logger, err := zapCfg.Build(option...)
	if err != nil {
		panic(err)
	}

	return logger
}
