package model

type Competitor struct {
	BaseModel
	ImageUrl string `json:"image_url"`
	Caption  string `json:"caption"`
	MatchID  uint64 `json:"match_id"`
}
