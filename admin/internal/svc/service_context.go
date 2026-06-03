// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"iot-platform-master/admin/internal/config"
	"iot-platform-master/models"
	"iot-platform-master/user/rpc/user_client"

	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user_client.User
	AuthUser *user_client.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user_client.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
