package service

import (
	"errors"
	"golang-server/database"
	"golang-server/model"
)

var MatchService = &matchService{repository: &database.MysqlMatchRepository{}}

type matchService struct {
	repository *database.MysqlMatchRepository
}

func (svc *matchService) FindAll(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	err := svc.repository.FindWithPageable(&matches, pageable)
	return &matches, err
}

func (svc *matchService) Save(match *model.Match) (*model.Match, error) {
	if !svc.isAvailable(match.Quota) || svc.isSuitablePayload(len(match.Competitors), match.Quota) {
		return nil, errors.New("등록할 자료 숫자가 유효하지 않습니다")
	}
	err := svc.repository.Save(match)
	return match, err
}

func (matchService) isAvailable(quota int) bool {
	condition := []int{16, 32, 64, 128}
	for _, c := range condition {
		if c == quota {
			return true
		}
	}
	return false
}

func (matchService) isSuitablePayload(sizeOfCompetitors int, quota int) bool {
	if sizeOfCompetitors == 0 {
		return true
	}
	return quota == sizeOfCompetitors
}

func (svc *matchService) FindOne(id uint64) (*model.Match, error) {
	match := new(model.Match)
	match.ID = id
	err := svc.repository.FindOne(match)
	return match, err
}
