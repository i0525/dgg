package validate

// 用户验证
type IndexRequest struct {
	Id   int64  `form:"id" valid:"required~id必须是整数"`
	Name string `form:"name" valid:"required~名字不能为空,runelength(1|10)~名字在1-10个字符之间"`
}