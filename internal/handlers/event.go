package handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"go.uber.org/zap"

	"feishuBot/internal/messages"
	"feishuBot/internal/services"
	"feishuBot/utils/logger"
)

func Webhook(c *gin.Context) {
	handler := dispatcher.NewEventDispatcher("J6P9r7s6Az54G64zt50eVhHwfWp0YVH3", "")
	handler = handler.OnP2MessageReceiveV1(func(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
		fmt.Println(larkcore.Prettify(event))
		fmt.Println(event.RequestId())

		resp, err := services.CallDeepSeekAPI(*event.Event.Message.Content)
		if err != nil {
			logger.Error("call deepseek api failed", zap.Error(err))
			return err
		}

		err = services.SendMessage(resp, *event.Event.Sender.SenderId.OpenId)
		if err != nil {
			logger.Error("send message failed", zap.Error(err))
			return err
		}

		return nil
	})

	req := messages.LarkEvent{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}

	logger.Debug("=====req", zap.Any("data", req))

	//
	//eventDetail, err := services.ParseEvent[messages.ReceiveMsgV2](&event)
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code":    -1,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//// Call DeepSeek API
	//deepSeekResp, err := services.CallDeepSeekAPI(eventDetail.Message.Content)
	//if err != nil {
	//	c.JSON(200, gin.H{
	//		"code":    -1,
	//		"message": err.Error(),
	//	})
	//	return
	//}
	//
	//err = services.SendMessage(deepSeekResp.Content, eventDetail.Sender.SenderId.OpenId)
}
