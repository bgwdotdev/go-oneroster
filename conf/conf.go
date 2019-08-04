package conf

import (
	"github.com/spf13/viper"
	errors "golang.org/x/xerrors"
	"os"
)

// Holds the config loaded from file
type Config struct {
	Port     int
	Database struct {
		DatabaseDriver string
		DataSourceName string
	}
}

type AuthConfig struct {
	Key    string
	KeyAlg string
}

// inits the JWT authorization config from environment variables
func (c *AuthConfig) Load() error {
	var ok bool
	c.Key, ok = os.LookupEnv("GOR_AUTH_KEY")
	if !ok {
		return errors.New("No JWT secret, set env GOR_AUTH_KEY")
	}
	c.KeyAlg, ok = os.LookupEnv("GOR_AUTH_KEYALG")
	if !ok {
		return errors.New("No JWT Key algrithm, set env GOR_AUTH_KEYALG; default HS256")
	}
	return nil
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
