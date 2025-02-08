package initializers

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
	"github.com/spf13/viper"
)

type Config struct {
	DBNAME     string `mapstructure:"DB_NAME"`
	DBHOST     string `mapstructure:"DB_HOST"`
	DBPORT     string `mapstructure:"DB_PORT"`
	SERVERADDR string `mapstructure:"SERVER_ADDR"`
	DBUSER     string `mapstructure:"DB_USER"`
	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
}

// GetConfig loads the configuration from .env
func GetConfig() Config {
	c, err := loadConfig()
	if err != nil {
		gologger.Fatal().Msgf("Error Loading Config: %s", err)
	}
	return c
}

func loadConfig() (config Config, err error) {
	viper.SetConfigFile(".env") //  Explicitly set .env file
	viper.SetConfigType("env")
	viper.AutomaticEnv() //Load system environment variables

	// Read the config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(" Could not read .env file. Falling back to system env variables.")
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	fmt.Println(" Final Loaded Config:")
	return config, nil
}
