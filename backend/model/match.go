package model

type Match struct {
	BaseModel
	MatchName string `json:"matchName" gorm:"column:match_name" validate:"required"`
	Quota     int    `json:"quota" gorm:"column:quota"`
	Available bool   `json:",omitempty" gorm:"column:available;default:0"`
}
