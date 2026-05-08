// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"iot-platform-master/admin/internal/config"
	"iot-platform-master/models"
"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
