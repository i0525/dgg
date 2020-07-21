package model

import "lidongheDemo/models"

// UsersModel 用户模型
type AdminsModel struct {
	models.BaseModel
	Id uint `gorm:"primary_key" json:"id"`
	Name string
}

// TableName 表名
func (u *AdminsModel) TableName() string {
	return "admins"
}

// UsersCreate 创建用户
//func UserCreate(u *UsersModel) (bool) {
//	rows := models.DB.Create(u).RowsAffected
//	if (rows < 1) {
//		return false
//	}
//	return true
//}
//
