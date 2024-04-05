package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func Login(db *gorm.DB, username, password string) (Account, error) {
	var account Account
	if err := db.Where("username = ? AND password = ?", username, password).First(&account).Error; err != nil {
		return Account{}, err
	}
	return account, nil
}

func Register(db *gorm.DB, account Account) (Account, error) {
	if err := db.Create(&account).Error; err != nil {
		return Account{}, err
	}
	return account, nil
}
