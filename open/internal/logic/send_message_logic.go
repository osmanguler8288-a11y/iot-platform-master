// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"iot-platform-master/device/types/device"
	"iot-platform-master/open/internal/svc"
	"iot-platform-master/open/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageRequest) (resp *types.SendMessageReply, err error) {
	if req.ProductKey == "" || req.DeviceKey == "" || req.Data == "" {
		return nil, errors.New("参数异常")
	}
	_, err = l.svcCtx.RpcDevice.SendMessage(l.ctx, &device.SendMessageRequest{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Data:       req.Data,
	})
	if err != nil {
		logx.Error("[ERROR] : ", err.Error())
		return nil, err
	}

	return &types.SendMessageReply{}, nil
}
