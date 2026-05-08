// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"iot-platform-master/admin/internal/svc"
	"iot-platform-master/admin/internal/types"
	"iot-platform-master/helper"
	"iot-platform-master/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeiviceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeiviceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeiviceListLogic {
	return &DeiviceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeiviceListLogic) DeiviceList(req *types.DeiviceListRequest) (resp *types.DeviceListReply, err error) {
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	var count int64
	resp = new(types.DeviceListReply)
	data := make([]*types.DeviceListBasic, 0)

	err = models.GetDeviceList(req.Name).Count(&count).Limit(req.Size).Offset(req.Page).Find(&data).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	resp.Count = count
	resp.List = data
	return
}
