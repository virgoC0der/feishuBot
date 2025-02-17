package handlers

import (
	"context"

	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"go.uber.org/zap"

	"feishuBot/internal/services"
	"feishuBot/utils/i18n"
	"feishuBot/utils/logger"
)

func ReceiveMsgHandler(ctx context.Context, event *larkim.P2MessageReceiveV1) error {
	logger.Debug("receive message", zap.String("event", larkcore.Prettify(event)))

	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("panic recovered", zap.Any("recover", r))
				_ = services.SendMessage(i18n.T("service_unavailable"), *event.Event.Sender.SenderId.OpenId)
			}
		}()

		openID := *event.Event.Sender.SenderId.OpenId
		msgContent := *event.Event.Message.Content

		handleError := func(err error, message string) {
			logger.Error(message, zap.Error(err))
			if sendErr := services.SendMessage(i18n.T("service_unavailable"), openID); sendErr != nil {
				logger.Error(i18n.T("send_error_failed"), zap.Error(sendErr))
			}
		}

		sendBotMessage := func(content string) error {
			return services.SendMessage(content, openID)
		}

		if msgContent == "" {
			return
		}

		if err := sendBotMessage(i18n.T("thinking")); err != nil {
			handleError(err, i18n.T("send_wait_message_failed"))
			return
		}

		resp, err := services.CallDeepSeekAPI(msgContent)
		if err != nil {
			handleError(err, i18n.T("api_call_failed"))
			return
		}

		if err := sendBotMessage(resp); err != nil {
			handleError(err, i18n.T("send_final_message_failed"))
		}
	}()

	return nil
}
