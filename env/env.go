package env

import (
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	DEBUG level = iota
	INFO
	ERROR
)

type level int

var (
	levels = map[string]level{
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"ERROR": ERROR,
	}
)

type Config struct {
	GRPCPort string
}

var Settings *Config

func Init() {
	logrus.SetOutput(os.Stdout)

	switch levels[viper.GetString("LOG_LEVEL")] {
	case DEBUG:
		logrus.SetLevel(logrus.DebugLevel)
	case INFO:
		logrus.SetLevel(logrus.InfoLevel)
	case ERROR:
		logrus.SetLevel(logrus.ErrorLevel)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetDefault("GRPC_PORT", 8080)

	Settings = &Config{
		GRPCPort: viper.GetString("GRPC_PORT"),
	}
}
