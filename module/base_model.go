package module

import "time"

// BaseModel 基模型定义
type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

//func (bm *BaseModel) BeforeSave() (err error) {
//	if bm.CreatedAt  != nil {
//		bm.CreatedAt = time.Now()
//	}
//	return
//}

func (bm *BaseModel) BeforeUpdate() (err error) {
	bm.UpdatedAt = time.Now()
	return
}

func (bm *BaseModel) AfterFind() (err error) {
	return
}
