package mysql

import (
	"github.com/jinzhu/gorm"
	"golang-server/model"
)

type (
	MysqlRepository struct{}

	Job func(db *gorm.DB) error
)

func (repo *MysqlRepository) FindOne(model interface{}) (err error) {
	return db.Find(model).Error
}

func (repo *MysqlRepository) FindAll(model interface{}, criteria interface{}) (err error) {
	return db.Where(criteria).Find(model).Error
}

func (repo *MysqlRepository) FindWithPageable(models interface{}, pageable *model.Pageable) (err error) {
	return db.Order(pageable.Order).Offset(pageable.Offset).Limit(pageable.Limit).Find(models).Error
}

func (repo *MysqlRepository) Save(model interface{}) (err error) {
	return repo.TransactionalJob(func(tx *gorm.DB) (err error) {
		return repo.Insert(tx, model)
	})
}

func (repo *MysqlRepository) Update(model interface{}, updated interface{}) (err error) {
	return db.Model(model).Updates(updated).Error
}

func (repo *MysqlRepository) Insert(tx *gorm.DB, model interface{}) (err error) {
	if !tx.NewRecord(model) {
		return err
	}
	return tx.Create(model).Error
}

func (repo *MysqlRepository) DefaultJob(fn Job) (err error) {
	return fn(db)
}

func (repo *MysqlRepository) TransactionalJob(fn Job) (err error) {
	tx := db.Begin()
	defer tx.Commit()
	if err = fn(tx); err != nil {
		tx.Rollback()
	}
	return
}
