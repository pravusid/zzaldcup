package model

import "time"

type BaseModel struct {
	ID        uint64     `json:"id" gorm:"primary_key,column:id" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:",omitempty" gorm:"column:deleted_at" sql:"index"`
}

func (model *BaseModel) BeforeCreate() (err error) {
	model.CreatedAt = time.Now()
	return
}

func (model *BaseModel) BeforeUpdate() (err error) {
	model.UpdatedAt = time.Now()
	return
}
