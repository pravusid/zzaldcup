package model

type Match struct {
	BaseModel
	MatchName   string       `json:"match_name"`
	Quota       int          `json:"quota"`
	Competitors []Competitor `gorm:"foreignkey:MatchID"`
}
