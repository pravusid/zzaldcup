package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang-server/model"
)

type (
	MysqlRepository struct{}

	Job func(db *gorm.DB) error
)

func (repo *MysqlRepository) FindOne(model interface{}) error {
	return db.Find(model).Error
}

func (repo *MysqlRepository) FindAll(model interface{}, criteria interface{}) error {
	return db.Where(criteria).Find(model).Error
}

func (repo *MysqlRepository) FindWithPageable(models interface{}, pageable *model.Pageable) error {
	return db.Order(pageable.Order).Offset(pageable.Offset).Limit(pageable.Limit).Find(models).Error
}

func (repo *MysqlRepository) Save(model interface{}) error {
	return repo.TransactionalJob(func(tx *gorm.DB) error {
		return repo.Insert(tx, model)
	})
}

func (repo *MysqlRepository) Insert(tx *gorm.DB, model interface{}) error {
	if !tx.NewRecord(model) {
		return errors.New("repository: record already exists")
	}
	return tx.Create(model).Error
}

func (repo *MysqlRepository) Update(model interface{}, updated interface{}) error {
	return db.Model(model).Updates(updated).Error
}

func (repo *MysqlRepository) Delete(model interface{}) error {
	return db.Delete(model).Error
}

func (repo *MysqlRepository) DefaultJob(fn Job) error {
	return fn(db)
}

func (repo *MysqlRepository) TransactionalJob(fn Job) error {
	tx := db.Begin()
	defer tx.Commit()
	if err := fn(tx); err != nil {
		tx.Rollback()
	}
	return nil
}
