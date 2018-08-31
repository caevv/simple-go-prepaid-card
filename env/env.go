package env

import (
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
	"os"
	"fmt"
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
	DBAddress string
	GRPCPort  string
}

var Settings *Config

func init() {
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
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASS", "postgres")
	viper.SetDefault("POSTGRES_HOST", "card-db")
	viper.SetDefault("POSTGRES_PORT", "5432")
	viper.SetDefault("POSTGRES_DB", "postgres")
	viper.SetDefault("POSTGRES_SSL", "false")

	Settings = &Config{
		DBAddress: address(),
		GRPCPort:  viper.GetString("GRPC_PORT"),
	}
}

func address() string {
	address := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=",
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASS"),
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"),
	)

	if !viper.GetBool("POSTGRES_SSL") {
		address += "disable"
	} else {
		address += "enable"
	}

	return address
}
