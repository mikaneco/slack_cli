package config

import (
	"os"
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
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(homedir, ".slk"))

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString("token")
}

func GetChannel() string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(homedir, ".slk"))

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return viper.GetString("channel_name")
}

func (c *Config) Save() error {

	homedir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Create the config directory if it doesn't exist
	configDir := filepath.Join(homedir, ".slk")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			return err
		}
	}

	configFilePath := filepath.Join(configDir, "config.json")

	// Set the configuration values
	viper.Set("channel_name", c.channel)
	viper.Set("token", c.token)

	// Write the configuration to disk
	viper.WriteConfigAs(configFilePath)

	return nil
}
