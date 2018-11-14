package service

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
)

var Match = &matchService{mysql: mysql.BaseMysqlRepository}

type matchService struct {
	mysql *mysql.MysqlRepository
}

func (svc *matchService) FindAll(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	pageable.Criteria = "available = true and private = false"
	err := svc.FindWithPageable(&matches, pageable)
	return &matches, err
}

func (svc *matchService) FindAllOfUser(pageable *model.Pageable) (*[]model.Match, error) {
	matches := make([]model.Match, pageable.Limit)
	err := svc.FindWithPageable(&matches, pageable)
	return &matches, err
}

func (svc *matchService) FindWithPageable(matches interface{}, pageable *model.Pageable) error {
	return svc.mysql.DefaultJob(func(db *gorm.DB) error {
		data := db.Where(pageable.Criteria).Order(pageable.Order)
		return data.Offset(pageable.Offset).Limit(pageable.Limit).Find(matches).Error
	})
}

func (svc *matchService) FindOne(id uint64) (*model.Match, error) {
	match := new(model.Match)
	match.ID = id
	err := svc.mysql.FindOne(match)
	return match, err
}

func (svc *matchService) FindOneByMatchName(matchName string) (*model.Match, error) {
	match := new(model.Match)
	return match, svc.mysql.DefaultJob(func(db *gorm.DB) error {
		return db.Where("match_name = ?", matchName).Find(match).Error
	})
}

func (svc *matchService) FindOneAndRelatedByMatchName(matchName string) (*model.Match, error) {
	match := new(model.Match)
	competitors := make([]model.Competitor, 16)
	err := svc.mysql.TransactionalJob(func(tx *gorm.DB) error {
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

func (svc *matchService) SaveOne(match *model.Match) error {
	if err := svc.isAvailable(match.Quota); err != nil {
		return err
	}
	return svc.mysql.Save(match)
}

func (svc *matchService) SavePrivateMatch(match *model.PrivateMatch) error {
	if err := svc.isAvailable(match.Match.Quota); err != nil {
		return err
	}
	return svc.mysql.Save(match)
}

func (svc *matchService) UpdateAvailability(tx *gorm.DB, criteria *model.Match, after int, before int) error {
	if err := tx.Find(&criteria).Error; err != nil {
		return err
	}
	if after > criteria.Quota {
		return errors.New("error: sufficient competitors")
	}
	anteChange := criteria.Available

	upward := after > before && after == criteria.Quota
	downward := after < before && before == criteria.Quota
	if upward || downward {
		if upward {
			criteria.Available = true
		} else if downward {
			criteria.Available = false
		}

		affected := tx.Model(criteria).Where(
			"available = ?", anteChange).Update("available", criteria.Available).RowsAffected
		if affected != 1 {
			return errors.New("error: simultaneous transactions")
		}
	}
	return nil
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
