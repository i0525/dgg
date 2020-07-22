package model

// UsersModel 用户模型
type AdminsModel struct {
	//module.BaseModel
	Id   uint `gorm:"primary_key" json:"id"`
	Name string
}

// TableName 表名
func (u *AdminsModel) TableName() string {
	return "admins"
}

// UsersCreate 创建用户
//func UserCreate(u *UsersModel) (bool) {
//	rows := module.DB.Create(u).RowsAffected
//	if (rows < 1) {
//		return false
//	}
//	return true
//}
//
