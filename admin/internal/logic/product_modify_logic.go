// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"iot-platform-master/admin/internal/svc"
	"iot-platform-master/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductModifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductModifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductModifyLogic {
	return &ProductModifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductModifyLogic) ProductModify(req *types.ProductModifyRequest) (resp *types.ProductModifyReply, err error) {
	// todo: add your logic here and delete this line

	return
}
