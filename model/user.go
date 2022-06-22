package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uint   `gorm:"not null;autoIncrement"`
	Email      string `gorm:"size:40;not null;unique"`
	Nickname   string `gorm:"column:name;size:40;not null"`
	Password   []byte `gorm:"column:password;not null"`
	InProgress string
}

func CreateUser(newUser *User) error {
	var err error
	newUser.Password, err = bcrypt.GenerateFromPassword(newUser.Password, bcrypt.DefaultCost) // encryption
	if err == nil {
		err = db.Create(newUser).Error
	}
	return err
}

func QueryUserByEmail(email string) (*User, error) {
	var dbUser User
	err := db.First(&dbUser, "Email = ?", email).Error
	return &dbUser, err
}

func QueryUserByUid(uid uint) (*User, error) {
	var dbUser User
	err := db.First(&dbUser, uid).Error
	return &dbUser, err
}

func UpdateUserPassword(uid uint, passwd []byte) error {
	password, err := bcrypt.GenerateFromPassword(passwd, bcrypt.DefaultCost) // encryption
	if err == nil {
		err = db.Model(&User{ID: uid}).Update("password", password).Error
	}
	return err
}

func UpdateUserNickname(uid uint, nickname string) error {
	err := db.Model(&User{ID: uid}).Update("name", nickname).Error
	return err
}
