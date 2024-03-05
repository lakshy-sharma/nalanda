package initializers

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Frontend Frontend `mapstructure:"frontend"`
	Backend  Backend  `mapstructure:"backend"`
	Database Database `mapstructure:"database"`
}
type Frontend struct {
}
type Backend struct {
	Port                   string `mapstructure:"port"`
	AccessTokenPrivateKey  string `mapstructure:"access_token_private_key"`
	AccessTokenPublicKey   string `mapstructure:"access_token_public_key"`
	RefreshTokenPrivateKey string `mapstructure:"refresh_token_private_key"`
	RefreshTokenPublicKey  string `mapstructure:"refresh_token_public_key"`
}
type Database struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DatabaseName string `mapstructure:"database_name"`
	ClientOrigin string `mapstructure:"client_origin"`
}

// This function takes a path to read the config from and then reads the configuration.
// The function is designed to handle toml based config files.
func LoadConfig(configPath string, configName string, configType string) (AppConfig, error) {
	// Create a new viper configurations reader.
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	// Parse the configuration into a structure.
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("Viper has failed to read the configurations: ", err)
		return AppConfig{}, err
	}
	// Unmarshal the configurations into the structure.
	var config AppConfig
	if err := v.Unmarshal(&config); err != nil {
		log.Fatal("Viper failed to unmarshal the configurations: ", err)
		return AppConfig{}, err
	}

	return config, nil
}
