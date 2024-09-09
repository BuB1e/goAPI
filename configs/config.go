
package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDatabase string `mapstructure:"DB_DATABASE"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost 	   string `mapstructure:"DB_HOST"`
	DBName 	   string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("❌ Error reading config file", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("❌ Unable to decode into struct", err)
	}

	return config
}