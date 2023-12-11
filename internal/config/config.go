package config

import "github.com/spf13/viper"

type Config struct {
	ConnString    string `mapstructure:"CONN_STRING"`
	DriverName    string `mapstructure:"DRIVER_NAME"`
	TelegramToken string `mapstructure:"TG_TOKEN"`
	Env           string `mapstructure:"ENV"`
}

func MustConfig() *Config {
	var config Config
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic("failed to read config: " + err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err.Error())
	}
	return &config
}
