package logic

import (
	"context"

	"errors"
	"iot-platform-master/device/internal/mqtt"
	"iot-platform-master/device/internal/svc"
	"iot-platform-master/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageReply, error) {
	if in.ProductKey == "" || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}
	topic := "/sys/" + in.ProductKey + "/" + in.DeviceKey + "/receive"
	if token := mqtt.MC.Publish(topic, 0, false, in.Data); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return &device.SendMessageReply{}, nil
}
