package database

import (
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
)

type MysqlCompetitorRepository struct {
	mysql.MysqlRepository
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
