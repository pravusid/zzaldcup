package service

import (
	"bytes"
	"errors"
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
	"io"
)

var Competitor = &competitorService{mysql: mysql.BaseMysqlRepository}

type competitorService struct {
	mysql *mysql.MysqlRepository
}

func (svc *competitorService) FindLatest(competitors *[]model.Competitor, criteria *model.Competitor) (*[]model.Competitor, error) {
	return competitors, svc.mysql.DefaultJob(func(db *gorm.DB) error {
		return db.Where(" match_id = ? AND id >= ?", criteria.MatchID, criteria.ID).Find(competitors).Error
	})
}

func (svc *competitorService) SaveOne(competitor *model.Competitor, match *model.Match) error {
	return svc.mysql.TransactionalJob(func(tx *gorm.DB) error {
		if err := svc.mysql.Insert(tx, competitor); err != nil {
			return err
		}

		var count int
		if err := tx.Model(model.Competitor{}).Where("match_id = ?", match.ID).Count(&count).Error; err != nil {
			return err
		}

		return Match.UpdateAvailability(tx, match, count, count-1)
	})
}

func (svc *competitorService) SaveFile(src io.Reader, ext string) (*model.ImagePath, error) {
	var path *model.ImagePath

	var buffer bytes.Buffer
	hash, err := FileService.HashingAndBuffering(&src, &buffer)
	if err != nil {
		return path, errors.New("error: file has fault")
	}

	var existence bool
	if path, existence = FileService.GenerateFilePath(hash, "image", ext); existence {
		return path, nil
	}

	return path, FileService.CreateFile(path, &buffer)
}

func (svc *competitorService) SaveAll(competitors []model.Competitor) error {
	return svc.mysql.TransactionalJob(func(tx *gorm.DB) error {
		for _, m := range competitors {
			if err := svc.mysql.Insert(tx, &m); err != nil {
				return err
			}
		}
		return nil
	})
}

func (svc *competitorService) UpdateOne(updated *model.Competitor) error {
	// TODO: user > match > competitor
	original := new(model.Competitor)
	original.ID = updated.ID
	if err := svc.mysql.FindOne(original); err != nil {
		return err
	}
	return svc.mysql.Update(updated, &model.Competitor{Caption: updated.Caption})
}

func (svc *competitorService) DeleteOne(competitor *model.Competitor) error {
	// TODO: user > match > competitor
	return svc.mysql.TransactionalJob(func(tx *gorm.DB) error {
		match, err := svc.competitorToMatch(tx, competitor)
		if err != nil {
			return err
		}

		if err := tx.Delete(competitor).Error; err != nil {
			return err
		}

		var count int
		if err = tx.Model(model.Competitor{}).Where("match_id = ?", match.ID).Count(&count).Error; err != nil {
			return err
		}

		return Match.UpdateAvailability(tx, match, count, count+1)
	})
}

func (svc *competitorService) competitorToMatch(tx *gorm.DB, competitor *model.Competitor) (*model.Match, error) {
	match := new(model.Match)
	if err := tx.Find(competitor).Error; err != nil {
		return match, err
	}
	match.ID = competitor.MatchID
	return match, tx.Find(match).Error
}
