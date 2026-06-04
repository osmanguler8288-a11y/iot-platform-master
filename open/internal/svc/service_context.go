// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"iot-platform-master/device/device_client"
	"iot-platform-master/open/internal/config"
	"iot-platform-master/user/rpc/user_client"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RpcUser   user_client.User
	RpcDevice device_client.Device
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcUser:   user_client.NewUser(zrpc.MustNewClient(c.RpcUser)),
		RpcDevice: device_client.NewDevice(zrpc.MustNewClient(c.RpcDevice)),
	}
}
