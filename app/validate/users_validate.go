package validate

// 用户验证
type Login struct {
	Id   int64  `form:"id" label:"用户名"  binding:"required"`
	Name   int64  `form:"name" binding:"required"`
}