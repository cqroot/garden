package config

import (
	"errors"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type configItem struct {
	Key   string
	Value any
}

type Config struct {
	configs       []configItem
	viperInstance *viper.Viper
}

func initDataDir() (string, error) {
	dataPath, err := xdg.DataFile("garden")
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(dataPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(dataPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return dataPath, nil
}

func New() (*Config, error) {
	dataPath, err := initDataDir()
	if err != nil {
		return nil, err
	}

	config := Config{
		configs: []configItem{
			{
				Key:   "bind_ip",
				Value: "127.0.0.1",
			},
			{
				Key:   "bind_port",
				Value: 8000,
			},
			{
				Key:   "log_level",
				Value: "Info",
			},
			{
				Key:   "log_with_caller",
				Value: false,
			},
			{
				Key:   "data_path",
				Value: dataPath,
			},
		},
		viperInstance: viper.New(),
	}

	config.viperInstance.SetTypeByDefaultValue(true)
	config.viperInstance.SetEnvPrefix("GARDEN")

	for _, conf := range config.configs {
		config.viperInstance.SetDefault(conf.Key, conf.Value)
		err := config.viperInstance.BindEnv(conf.Key)
		if err != nil {
			return nil, err
		}
	}

	return &config, nil
}

func (c *Config) BindIp() string {
	return c.viperInstance.GetString("bind_ip")
}

func (c *Config) BindPort() int {
	return c.viperInstance.GetInt("bind_port")
}

func (c *Config) LogLevel() string {
	return c.viperInstance.GetString("log_level")
}

func (c *Config) LogWithCaller() bool {
	return c.viperInstance.GetBool("log_with_caller")
}

func (c *Config) DataPath() string {
	return c.viperInstance.GetString("data_path")
}
