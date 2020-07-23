package validate

// 用户验证
type Login struct {
	Id   int64  `form:"id" binding:"required"`
	Name   int64  `form:"name" binding:"required"`
}