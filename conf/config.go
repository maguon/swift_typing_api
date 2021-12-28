package conf

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Schema struct {
	Env          string `mapstructure:"env"`
	DefaultLimit int    `mapstructure:"default_limit"`
	MaxLimit     int    `mapstructure:"max_limit"`

	Log struct {
		LogLevel int    `mapstructure:"log_level"`
		LogCount int    `mapstructure:"log_count"`
		LogPath  string `mapstructure:"log_path"`
		LogName  string `mapstructure:"log_name"`
	}

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Name     string `mapstructure:"name"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Env      string `mapstructure:"env"`
		SSLMode  string `mapstructure:"sslmode"`
	} `mapstructure:"database"`

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		Database int    `mapstructure:"database"`
	} `mapstructure:"redis"`
}

var Config Schema

func init() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./conf") // Look for config in current directory

	config.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	config.AutomaticEnv()

	err := config.ReadInConfig() // Find and read the config file
	if err != nil {              // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	err = config.Unmarshal(&Config)
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	// fmt.Printf("Current Config: %+v", Config)
}
