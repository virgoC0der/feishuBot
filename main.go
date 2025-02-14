package main

import (
	"github.com/gin-gonic/gin"
	sdkginext "github.com/larksuite/oapi-sdk-gin"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"go.uber.org/zap"

	"feishuBot/internal/conf"
	"feishuBot/internal/handlers"
	"feishuBot/internal/services"
	"feishuBot/utils/logger"
)

//func main() {
//
//	resp, err := services.CallDeepSeekAPI("通过go来实现一个飞书机器人的后端，通过使用webhook，接收到消息后，请求deepseek api，将deepseek返回的content回复到飞书")
//	if err != nil {
//		logger.Error("call deepseek api failed", zap.Error(err))
//		return
//	}
//
//	logger.Debug("=====llm resp", zap.String("resp", resp))
//	err = services.SendMessage(resp, "ou_8706df349d1affe1868684dcc9f3580a")
//	if err != nil {
//		logger.Error("send message failed", zap.Error(err))
//		return
//	}
//
//}

func main() {
	if err := conf.InitConf(); err != nil {
		logger.Fatal("init config failed", zap.Error(err))
	}

	services.InitLark()
	services.InitOpenAI()

	handler := dispatcher.NewEventDispatcher("J6P9r7s6Az54G64zt50eVhHwfWp0YVH3", "")
	handler = handler.OnP2MessageReceiveV1(handlers.ReceiveMsgHandler)

	g := gin.Default()

	api := g.Group("/api/v1")
	{
		api.POST("/feishu/webhook", sdkginext.NewEventHandlerFunc(handler))
	}

	err := g.Run(":8081")
	if err != nil {
		panic(err)
	}
}
