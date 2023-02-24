package config

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	channel string
	token   string
}

func New() *Config {
	return &Config{}
}

func (c *Config) SetChannel(channel string) {
	c.channel = channel
}

func (c *Config) SetToken(token string) {
	c.token = token
}

func (c *Config) GetChannel() string {
	return c.channel
}

func (c *Config) GetToken() string {
	return c.token
}

func GetToken() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString("token")
}

func GetChannel() string {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString("channel_name")
}

func (c *Config) Save() error {
	configFilePath := filepath.Join("./", "config", "config.json")

	// Set the configuration values
	viper.Set("channel_name", c.channel)
	viper.Set("token", c.token)

	// Save the configuration to the file
	err := viper.WriteConfigAs(configFilePath)
	if err != nil {
		return err
	}

	return nil
}
