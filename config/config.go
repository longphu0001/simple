package config

import "github.com/spf13/viper"

// Config stores configuration variables.
type Config struct {
	Database     *Database
	HTTP         *HTTP
	Release      *Release
}

// New returns a new config instance.
func New() (*Config, error) {
	cfg := &Config{}

	viper.New()
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath("$GOPATH/src/go-learn/projects/simple/config/.")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
