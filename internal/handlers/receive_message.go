package handlers

import (
	"context"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"go.uber.org/zap"

	"feishuBot/internal/services"
	"feishuBot/utils/logger"
)

func ReceiveMsgHandler(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	logger.Debug("receive message", zap.String("event", larkcore.Prettify(event)))

	go func() {
		if err := services.SendMessage("正在思考中...", *event.Event.Sender.SenderId.OpenId); err != nil {
			logger.Error("send message failed", zap.Error(err))
			return
		}

		resp, err := services.CallDeepSeekAPI(*event.Event.Message.Content)
		if err != nil {
			logger.Error("call deepseek api failed", zap.Error(err))
			return
		}

		err = services.SendMessage(resp, *event.Event.Sender.SenderId.OpenId)
		if err != nil {
			logger.Error("send message failed", zap.Error(err))
			return
		}
	}()

	return nil
}
