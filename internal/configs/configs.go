package configs

import "github.com/spf13/viper"

func init() {
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvPrefix("TODO")

	viper.SetDefault("bind_ip", "127.0.0.1")
	viper.SetDefault("bind_port", 8000)
	viper.SetDefault("log_level", "Debug")

	viper.BindEnv("bind_ip")
	viper.BindEnv("bind_port")
	viper.BindEnv("log_level")
}

func BindIp() string {
	return viper.GetString("bind_ip")
}

func BindPort() int {
	return viper.GetInt("bind_port")
}

func LogLevel() string {
	return viper.GetString("log_level")
}
