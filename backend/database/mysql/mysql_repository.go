package mysql

import (
	"github.com/jinzhu/gorm"
	"golang-server/model"
)

type MysqlRepository struct{}

func (repo *MysqlRepository) Close() {
	db.Close()
}

func (repo *MysqlRepository) FindOne(model interface{}) (err error) {
	if err = db.Where(model).Find(model).Error; err != nil {
		return err
	}
	return
}

func (repo *MysqlRepository) FindAll(model interface{}, condition interface{}) (err error) {
	if err = db.Where(condition).Find(model).Error; err != nil {
		return err
	}
	return
}

func (repo *MysqlRepository) FindWithPageable(models interface{}, pageable *model.Pageable) (err error) {
	db.Order(pageable.Order).Offset(pageable.Offset).Limit(pageable.Limit).Find(models)
	return
}

func (repo *MysqlRepository) Save(model interface{}) (err error) {
	return repo.TransactionalJob(func(tx *gorm.DB) (err error) {
		return repo.Insert(tx, model)
	})
}

func (repo *MysqlRepository) Insert(tx *gorm.DB, model interface{}) (err error) {
	if !tx.NewRecord(model) {
		return err
	}
	if err = tx.Create(model).Error; err != nil {
		return err
	}
	return
}

func (repo *MysqlRepository) TransactionalJob(fn func(transaction *gorm.DB) error) (err error) {
	tx := db.Begin()
	defer tx.Commit()
	if err = fn(tx); err != nil {
		tx.Rollback()
	}
	return err
}
