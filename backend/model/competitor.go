package model

type Competitor struct {
	BaseModel
	ImageUrl string
	Caption  string
	MatchID  uint64
}
