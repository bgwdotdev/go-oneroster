package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// Holds the config loaded from file
type Config struct {
	Port     int
	Database []struct {
		DatabaseDriver string
		DataSourceName string
	}
}

// Invokes Viper to read the configuration file
func readConf() (Config, error) {
	var c Config

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.SetConfigType("hcl")
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	fmt.Println("Loaded conf.hcl")

	err = viper.Unmarshal(&c)
	return c, err
}

func main() {
	readConf()
}
