package infrastructure

import (
	"github.com/spf13/viper"
	"github.com/vnnyx/gas-api/exception"
)

type Config struct {
	AppPort                string `mapstructure:"APP_PORT"`
	MysqlHostSlave         string `mapstructure:"MYSQL_HOST_SLAVE"`
	MysqlPoolMin           int    `mapstructure:"MYSQL_POOL_MIN"`
	MysqlPoolMax           int    `mapstructure:"MYSQL_POOL_MAX"`
	MysqlIdleMax           int    `mapstructure:"MYSQL_IDLE_MAX"`
	MysqlMaxIdleTimeMinute int    `mapstructure:"MYSQL_MAX_IDLE_TIME_MINUTE"`
	MysqlMaxLifeTimeMinute int    `mapstructure:"MYSQL_MAX_LIFE_TIME_MINUTE"`
}

func NewConfig(configName string) *Config {
	config := &Config{}
	viper.AddConfigPath(".")
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	exception.PanicIfNeeded(err)
	err = viper.Unmarshal(&config)
	exception.PanicIfNeeded(err)
	return config
}
