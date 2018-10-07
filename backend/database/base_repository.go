package database

import (
	"github.com/jinzhu/gorm"
	"golang-server/model"
)

func FindOne(model interface{}) (err error) {
	if err = db.Where(model).Find(model).Error; err != nil {
		return err
	}
	return
}

func FindAll(models interface{}, pageable *model.Pageable) (err error) {
	db.Order(pageable.Order).Offset(pageable.Offset).Limit(pageable.Limit).Find(models)
	return
}

func Save(model interface{}) (err error) {
	if !db.NewRecord(model) {
		return err
	}
	return transactionalJob(func(tx *gorm.DB) (err error) {
		if err = tx.Create(model).Error; err != nil {
			return err
		}
		return
	})
}

func transactionalJob(fn func(transaction *gorm.DB) error) (err error) {
	tx := db.Begin()
	defer tx.Commit()
	if err = fn(tx); err != nil {
		tx.Rollback()
	}
	return err
}
