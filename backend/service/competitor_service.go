package service

import (
	"golang-server/database"
	"golang-server/model"
)

var CompetitorService = &competitorService{repository: &database.MysqlCompetitorRepository{}}

type competitorService struct {
	repository *database.MysqlCompetitorRepository
}

func (svc *competitorService) FindAll(competitors *[]model.Competitor, condition uint64) (*[]model.Competitor, error) {
	match := new(model.Match)
	match.ID = condition
	err := svc.repository.FindAll(competitors, match)
	return competitors, err
}

func (svc *competitorService) Save(competitors *[]model.Competitor) (*[]model.Competitor, error) {
	err := svc.repository.SaveAll(*competitors)
	return competitors, err
}

func (svc *competitorService) FindOne(id uint64) (*model.Competitor, error) {
	competitor := new(model.Competitor)
	competitor.ID = id
	err := svc.repository.FindOne(&competitor)
	return competitor, err
}
