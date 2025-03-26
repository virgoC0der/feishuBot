# feishuBot

## 📖 项目简介

feishuBot 是一个接入飞书机器人的后端服务，可以将飞书机器人与GPT等大语言模型连接，实现智能对话功能。通过简单的配置，你可以快速搭建自己的智能对话机器人，提升工作效率。

## ✨ 特性

- 🤖 接入飞书机器人，作为GPT机器人的后端服务
- 🔌 支持多种大语言模型接入
- ⚙️ 简单易用的YAML配置
- 🚀 高性能，低延迟

## 🔧 安装

### 前置条件

- Go 环境 (推荐 Go 1.21+)
- 飞书开发者账号
- GPT或其他大语言模型的API密钥

### 安装步骤

#### 本地运行

```bash
# 克隆仓库
git clone https://github.com/virgoC0der/feishuBot.git
cd feishuBot

# 安装依赖
go mod tidy

# 配置环境变量（可选，默认使用中文）
export LANG=zh-CN  # 或 en-US

# 启动服务
go run main.go
```

#### 使用Docker

```bash
# 拉取镜像
docker pull ghcr.io/virgoc0der/feishubot:main
# 或使用 GitHub 容器注册表
# docker pull ghcr.io/用户名/feishubot:latest

# 运行容器
docker run -p 8081:8081 \
  -e LANG=zh-CN \
  -v /path/to/config:/root/config \
  ghcr.io/virgoc0der/feishubot:main
```

## ⚙️ 配置

在项目根目录创建 `config.yaml` 文件，并按以下格式配置：

```yaml
feishu:
  app_id: YOUR_APP_ID          # 飞书应用ID
  app_secret: YOUR_APP_SECRET  # 飞书应用密钥
  verify_token: YOUR_VERIFY_TOKEN  # 飞书事件订阅的验证令牌
llm:
  model: YOUR_MODEL            # 使用的模型，如 gpt-3.5-turbo
  base_url: YOUR_API_BASE_URL  # API基础URL
  api_key: YOUR_API_KEY        # API密钥
```

## 📝 飞书机器人配置

1. 在[飞书开发者平台](https://open.feishu.cn/)创建应用
2. 开启机器人功能
3. 配置消息卡片请求网址为你的服务地址
4. 添加事件订阅，URL为你的服务地址/webhook
5. 订阅必要的事件（如接收消息等）

## 📄 许可证

[MIT License](LICENSE)

## 🙏 贡献

欢迎提交问题和PR，一起改进这个项目！

## 📞 联系方式

如有问题，请通过Issues或Pull Requests联系我们。
