package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang-server/database"
	"golang-server/model"
)

var MatchService = &matchService{repository: &database.MysqlMatchRepository{}}

type matchService struct {
	repository *database.MysqlMatchRepository
}

func (svc *matchService) FindAll(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	pageable.Criteria = "available = true and private = false"
	err := svc.repository.FindWithPageable(&matches, pageable)
	return &matches, err
}

func (svc *matchService) FindUserMatches(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	err := svc.repository.FindWithPageable(&matches, pageable)
	return &matches, err
}

func (svc *matchService) FindOne(id uint64) (*model.Match, error) {
	match := new(model.Match)
	match.ID = id
	err := svc.repository.FindOne(match)
	return match, err
}

func (svc *matchService) FindOneByMatchName(matchName string) (*model.Match, error) {
	match := new(model.Match)
	match.MatchName = matchName
	err := svc.repository.FindOne(match)
	return match, err
}

func (svc *matchService) FindOneAndRelatedByMatchName(matchName string) (*model.Match, error) {
	match := new(model.Match)
	competitors := make([]model.Competitor, 16)
	err := svc.repository.TransactionalJob(func(tx *gorm.DB) error {
		if err := tx.Where("match_name = ?", matchName).Find(match).Error; err != nil {
			return err
		}
		if err := tx.Model(match).Related(&competitors).Error; err != nil {
			return err
		}
		return nil
	})
	match.Competitors = competitors
	return match, err
}

func (svc *matchService) Save(match *model.Match) error {
	if err := svc.isAvailable(match.Quota); err != nil {
		return err
	}
	return svc.repository.Save(match)
}

func (svc *matchService) SavePrivate(match *model.PrivateMatch) error {
	if err := svc.isAvailable(match.Match.Quota); err != nil {
		return err
	}
	return svc.repository.Save(match)
}

func (matchService) isAvailable(quota int) error {
	condition := []int{16, 32, 64, 128}
	for _, c := range condition {
		if c == quota {
			return nil
		}
	}
	return errors.New("error: unavailable quota")
}

func (matchService) isSuitablePayload(sizeOfCompetitors int, quota int) bool {
	return quota == sizeOfCompetitors
}
