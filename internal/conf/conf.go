package conf

import (
	"feishuBot/utils/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type LarkConf struct {
	AppId       string `mapstructure:"app_id"`
	AppSecret   string `mapstructure:"app_secret"`
	VerifyToken string `mapstructure:"verify_token"`
}

type LLMConf struct {
	Model   string `mapstructure:"model"`
	BaseUrl string `mapstructure:"base_url"`
	ApiKey  string `mapstructure:"api_key"`
}

type Config struct {
	Lark LarkConf `mapstructure:"feishu"`
	LLM  LLMConf  `mapstructure:"llm"`
}

var (
	GConfig Config
)

func InitConf() error {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/root/config/")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error("read config failed", zap.Error(err))
		return err
	}

	if err := viper.Unmarshal(&GConfig); err != nil {
		logger.Error("unmarshal config failed", zap.Error(err))
		return err
	}

	logger.Debug("init config success", zap.Any("config", GConfig))
	return nil
}
