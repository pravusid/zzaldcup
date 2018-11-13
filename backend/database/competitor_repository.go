package database

import (
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
)

type MysqlCompetitorRepository struct {
	mysql.MysqlRepository
}

func (repo *MysqlCompetitorRepository) FindWithCursor(models interface{}, criteria *model.Competitor) (err error) {
	return repo.DefaultJob(func(db *gorm.DB) error {
		return db.Where(" match_id = ? AND id >= ?", criteria.MatchID, criteria.ID).Find(models).Error
	})
}

func (repo *MysqlCompetitorRepository) Count(count *int, criteria *model.Competitor) (err error) {
	return repo.DefaultJob(func(db *gorm.DB) error {
		return db.Model(&model.Competitor{}).Where(criteria).Count(count).Error
	})
}

func (repo *MysqlCompetitorRepository) SaveAll(models []model.Competitor) (err error) {
	return repo.TransactionalJob(func(tx *gorm.DB) (err error) {
		for _, m := range models {
			if err := repo.Insert(tx, &m); err != nil {
				return err
			}
		}
		return
	})
}
