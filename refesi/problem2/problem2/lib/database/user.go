package database

import (
	"fmt"
	"problem2/model"

	"gorm.io/gorm"
)

func (gorm.DB) UserGetAll() ([]*model.User, error) {

	var users []*model.User
	err := db.Tx.Find(&users).Error
	if err != nil {
		db.Tx.Rollback()
		panic(err)
	}

	return users, nil
}

func (d *db) UserCreate(input model.User) (*model.User, error) {

	err := d.Tx.Model(&model.User{}).Create(&input).Error
	if err != nil {
		d.Tx.Rollback()
		panic(err)
	}

	return &input, nil
}

func (d *db) UserGetByID(id int) (*model.User, error) {

	var user model.User
	err := d.Tx.Model(&model.User{}).Where("id = ?", id).Take(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			d.Tx.Rollback()
			return nil, fmt.Errorf("user not available")
		}
	}

	return &user, nil
}

func (d *db) UserGetByEmail(email string) (*model.User, error) {

	var user model.User
	err := d.Tx.Model(&model.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			d.Tx.Rollback()
			panic(err)
		} else {
			return nil, fmt.Errorf("user not available")
		}
	}

	return &user, nil
}
