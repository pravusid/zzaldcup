package database

import (
	"github.com/jinzhu/gorm"
	"golang-server/database/mysql"
	"golang-server/model"
)

var MysqlRepository = &baseRepository{db: mysql.DB()}

func CloseAll() {
	MysqlRepository.close()
}

type baseRepository struct {
	db *gorm.DB
}

func (repo *baseRepository) FindOne(model interface{}) (err error) {
	if err = repo.db.Where(model).Find(model).Error; err != nil {
		return err
	}
	return
}

func (repo *baseRepository) FindAll(models interface{}, pageable *model.Pageable) (err error) {
	repo.db.Order(pageable.Order).Offset(pageable.Offset).Limit(pageable.Limit).Find(models)
	return
}

func (repo *baseRepository) Save(model interface{}) (err error) {
	if !repo.db.NewRecord(model) {
		return err
	}
	return repo.transactionalJob(func(tx *gorm.DB) (err error) {
		if err = tx.Create(model).Error; err != nil {
			return err
		}
		return
	})
}

func (repo *baseRepository) transactionalJob(fn func(transaction *gorm.DB) error) (err error) {
	tx := repo.db.Begin()
	defer tx.Commit()
	if err = fn(tx); err != nil {
		tx.Rollback()
	}
	return err
}

func (repo *baseRepository) close() {
	repo.db.Close()
}
