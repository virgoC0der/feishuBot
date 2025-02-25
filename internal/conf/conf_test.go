package conf

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestInitConf(t *testing.T) {
	// 调整测试配置路径到测试用的 config 目录（或你自己的路径）
	viper.AddConfigPath(filepath.Join("..", "..", "testdata"))
	// 假设你在 testdata 目录下准备了名为 app.yaml 的测试配置文件

	err := InitConf()
	assert.NoError(t, err, "InitConf 应该成功读取并解析配置")

	// 检查全局 GConfig 是否已经赋值
	assert.NotEmpty(t, GConfig.Lark.AppId, "测试配置文件中的 Lark.AppId 不应为空")
	assert.NotEmpty(t, GConfig.Lark.AppSecret, "测试配置文件中的 Lark.AppSecret 不应为空")
	assert.NotEmpty(t, GConfig.Lark.VerifyToken, "测试配置文件中的 Lark.VerifyToken 不应为空")

	assert.NotEmpty(t, GConfig.LLM.Model, "测试配置文件中的 LLM.Model 不应为空")
	assert.NotEmpty(t, GConfig.LLM.BaseUrl, "测试配置文件中的 LLM.BaseUrl 不应为空")
	assert.NotEmpty(t, GConfig.LLM.ApiKey, "测试配置文件中的 LLM.ApiKey 不应为空")
}
