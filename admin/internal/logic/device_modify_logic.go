// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"iot-platform-master/admin/internal/svc"
	"iot-platform-master/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceModifyLogic {
	return &DeviceModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceModifyLogic) DeviceModify(req *types.DeviceModifyRequest) (resp *types.DeviceModifyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
