package model

type Match struct {
	BaseModel
	MatchName string `json:"matchName" gorm:"column:match_name"`
	Quota     int    `json:"quota" gorm:"column:quota"`
}
