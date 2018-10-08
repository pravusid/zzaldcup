package service

import (
	"golang-server/database"
	"golang-server/model"
)

type matchService struct{}

var MatchService = &matchService{}

func (svc *matchService) FindAll(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	err := database.MysqlRepository.FindAll(&matches, pageable)
	return &matches, err
}

func (svc *matchService) Save(match *model.Match) (*model.Match, error) {
	err := database.MysqlRepository.Save(match)
	return match, err
}

func (svc *matchService) FindOne(id uint64) (*model.Match, error) {
	match := new(model.Match)
	match.ID = id
	err := database.MysqlRepository.FindOne(&match)
	return match, err
}
