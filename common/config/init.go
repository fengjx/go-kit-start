package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"

	"github.com/fengjx/go-kit-start/common/env"
)

var conf *AppConfig

func init() {
	viperConfig := viper.New()
	loadBase(viperConfig)
	mergeConfig(viperConfig)
	c := Config{}
	err := viperConfig.Unmarshal(&c, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "yaml"
	})
	if err != nil {
		panic(err)
	}
	conf = &AppConfig{
		Viper:  viperConfig,
		Config: c,
	}
}

func loadBase(viperConfig *viper.Viper) {
	configName := "conf/app.yaml"
	viperConfig.SetConfigFile(configName)
	err := viperConfig.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func mergeConfig(viperConfig *viper.Viper) {
	var configFile string
	envConfigPath := os.Getenv("APP_CONFIG")
	if envConfigPath != "" {
		configFile = envConfigPath
	}
	if configFile == "" && len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	if configFile == "" {
		appEnv := env.GetEnv()
		configFile = path.Join("conf", fmt.Sprintf("app-%s.yaml", appEnv))
	}
	viperConfig.SetConfigFile(configFile)
	err := viperConfig.MergeInConfig()
	if err != nil {
		panic(err)
	}
}

// GetConfig 返回项目配置
func GetConfig() *AppConfig {
	return conf
}
