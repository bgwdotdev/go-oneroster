package conf

import (
	"github.com/spf13/viper"
)

// Holds the config loaded from file
type Config struct {
	Port     int
	Database struct {
		DatabaseDriver string
		DataSourceName string
	}
}

// Invokes Viper to read the configuration file
func Read() (Config, error) {
	var c Config

	viper.SetConfigName("conf")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = viper.Unmarshal(&c)
	return c, err
}
