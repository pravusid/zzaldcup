package model

type Match struct {
	BaseModel
	MatchName   string       `json:"matchName" gorm:"unique;column:match_name" validate:"required"`
	Quota       int          `json:"quota" gorm:"column:quota"`
	Private     bool         `json:"private" gorm:"column:private"`
	Available   bool         `json:"available" gorm:"column:available;default:0"`
	Competitors []Competitor `json:"competitors" gorm:"foreignkey:MatchID"`
}
