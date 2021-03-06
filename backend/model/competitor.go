package model

type Competitor struct {
	BaseModel
	ImageUrl string `json:"imageUrl" gorm:"column:image_url"`
	Caption  string `json:"caption" gorm:"column:caption"`
	MatchID  uint64 `json:"-" gorm:"column:match_id"`
}
