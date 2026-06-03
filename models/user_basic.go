package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity;type:varchar(50);" json:"identity"`
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	Password string `gorm:"column:password;type:varchar(50);" json:"password"`
	//这里的gorm就是在备注，然后column就是在表示列的名字，本质上就是在进行一个命名，对于不同的列在命名
	//中间的type就是在给数据库列类型，最后这个json就是在其序列化时的字段名什么的
	//gorm.model是怎么个情况呢？就是内部镶嵌了一部分的gorm需要的数据，包括id，增删查改时间什么的
}
func (table UserBasic) TableName() string {
	return "user_basic"
}