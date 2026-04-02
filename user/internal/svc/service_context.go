// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"gitee/getcharzp/iot-platform/models"
	"iot-platform-master/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:    models.DB, // 暂时注释掉，models包未导入
	}
}
