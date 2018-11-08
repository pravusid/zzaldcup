package database

import (
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
)

type MysqlMatchRepository struct {
	mysql.MysqlRepository
}

func (repo *MysqlMatchRepository) FindWithPageable(models interface{}, pageable *model.Pageable) (err error) {
	return repo.DefaultJob(func(db *gorm.DB) error {
		data := db.Where("available = true and private = false").Order(pageable.Order)
		return data.Offset(pageable.Offset).Limit(pageable.Limit).Find(models).Error
	})
}
