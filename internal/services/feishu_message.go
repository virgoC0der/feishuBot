package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"go.uber.org/zap"

	"feishuBot/internal/conf"
	"feishuBot/utils/logger"
)

var (
	feishuCli *lark.Client
)

func InitLark() {
	feishuCli = lark.NewClient(conf.GConfig.Lark.AppId, conf.GConfig.Lark.AppSecret)
}

func SendMessage(msg string, receiveId string) error {
	contentJson, err := json.Marshal(map[string]string{"text": msg})
	if err != nil {
		logger.Error("send message failed", zap.Error(err))
		return err
	}

	req := larkim.NewCreateMessageReqBuilder().
		ReceiveIdType("open_id").
		Body(larkim.NewCreateMessageReqBodyBuilder().
			ReceiveId(receiveId).
			MsgType("text").
			Content(string(contentJson)).
			Build()).
		Build()

	resp, err := feishuCli.Im.V1.Message.Create(context.Background(), req)
	if err != nil {
		logger.Error("send message failed", zap.Error(err))
		return err
	}

	if !resp.Success() {
		logger.Error("send message failed", zap.Any("resp", resp))
		return fmt.Errorf("send message failed")
	}

	return nil
}
