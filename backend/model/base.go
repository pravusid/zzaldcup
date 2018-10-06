package model

import "time"

type BaseModel struct {
	ID        uint64     `gorm:"primary_key,column:id" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time  `gorm:"column:created_at" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `gorm:"column:updated_at" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `sql:"index"`
}
