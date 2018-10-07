package model

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"primary_key,column:id" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:",omitempty" sql:"index"`
}

func (model *BaseModel) BeforeCreate() (err error) {
	model.CreatedAt = time.Now()
	return
}

func (model *BaseModel) BeforeUpdate() (err error) {
	model.UpdatedAt = time.Now()
	return
}
