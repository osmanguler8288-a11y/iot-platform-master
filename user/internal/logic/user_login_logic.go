// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"

	"iot-platform-master/helper"
	"iot-platform-master/models"
	"iot-platform-master/user/internal/svc"
	"iot-platform-master/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginRequest) (resp *types.UsrLoginReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.UsrLoginReply)
	ub := new(models.UserBasic)
	err = l.svcCtx.DB.Where("name=? AND password=?",
		req.Username, helper.Md5(req.Password)).First(ub).Error
	if err != nil {
		logx.Error("DB error:", err)
		err = errors.New("用户名或者密码不正确")
		return
	}
	token, err := helper.GenerateToken(ub.ID, ub.Identity, ub.Name, 3600*24*30)
	if err != nil {
		logx.Error("GenerateToken error:", err)
		err = errors.New("用户名或者密码不正确、token生成失败")
		return
	}
	resp.Token = token
	return
}
