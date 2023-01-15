package app

import (
	"github.com/adrg/xdg"
	"github.com/spf13/viper"
)

type configItem struct {
	Key   string
	Value any
}

type ConfigProvider struct {
	configs       []configItem
	viperInstance *viper.Viper
}

var configProvider ConfigProvider

func InitConfig() error {
	dataPath, err := xdg.DataFile("garden")
	if err != nil {
		return err
	}

	configProvider = ConfigProvider{
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
				Key:   "data_path",
				Value: dataPath,
			},
		},
		viperInstance: viper.New(),
	}

	configProvider.viperInstance.SetTypeByDefaultValue(true)
	configProvider.viperInstance.SetEnvPrefix("GARDEN")

	for _, conf := range configProvider.configs {
		configProvider.viperInstance.SetDefault(conf.Key, conf.Value)
		err := configProvider.viperInstance.BindEnv(conf.Key)
		if err != nil {
			return err
		}
	}

	return nil
}

func Config() *ConfigProvider {
	return &configProvider
}

func (c *ConfigProvider) BindIp() string {
	return c.viperInstance.GetString("bind_ip")
}

func (c *ConfigProvider) BindPort() int {
	return c.viperInstance.GetInt("bind_port")
}

func (c *ConfigProvider) LogLevel() string {
	return c.viperInstance.GetString("log_level")
}

func (c *ConfigProvider) DataPath() string {
	return c.viperInstance.GetString("data_path")
}
