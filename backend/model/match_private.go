package model

type PrivateMatch struct {
	BaseModel
	Match   Match `gorm:"foreignkey:MatchID"`
	MatchID uint64
	UUID    string `json:"uuid" gorm:"column:uuid"`
}
