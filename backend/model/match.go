package model

type Match struct {
	BaseModel
	MatchName   string
	Competitors []Competitor `gorm:"foreignkey:MatchID"`
}
