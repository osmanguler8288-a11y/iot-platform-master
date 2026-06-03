package logic

import (
	"context"
	"errors"

	"iot-platform-master/helper"
	"iot-platform-master/user/rpc/internal/svc"
	"iot-platform-master/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	if in.Token == "" {
		return nil, errors.New("必填参数不能为空")
	}
	userClaim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	resp := new(user.UserAuthReply)
	resp.Identity = userClaim.Identity
	resp.Id = uint64(userClaim.Id)
	resp.Extend = map[string]string{
		"name": userClaim.Name,
	}
	return resp, nil
}
