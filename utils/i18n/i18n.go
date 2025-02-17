package i18n

import (
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
)

func init() {
	// 从环境变量获取语言设置，默认中文
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = "zh-CN"
	}

	bundle = i18n.NewBundle(language.MustParse(lang))
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// 加载语言文件
	bundle.MustLoadMessageFile("locales/en-US.yaml")
	bundle.MustLoadMessageFile("locales/zh-CN.yaml")

	// 创建本地化器
	localizer = i18n.NewLocalizer(bundle, lang)
}

// 根据语言创建新的localizer
func NewLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}

// 带语言参数的翻译函数
func T(messageID string, args ...interface{}) string {
	config := &i18n.LocalizeConfig{
		MessageID: messageID,
	}

	if len(args) > 0 {
		config.TemplateData = args[0]
	}

	str, err := localizer.Localize(config)
	if err != nil {
		return messageID // 返回消息ID作为fallback
	}
	return str
}
