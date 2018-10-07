package model

type Match struct {
	BaseModel
	MatchName   string       `json:"match_name"`
	Competitors []Competitor `gorm:"foreignkey:MatchID"`
}
