package schemas

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        uint           `gorm: "primaryKey"`
	Username  string         `gorm: "type: text"`
	Password  string         `gorm: "type: text"`
	Email     string         `gorm: "type: text"`
	CreatedAt time.Time      `gorm: "autoCreateTime"`
	UpdatedAt time.Time      `gorm: "autoUpdatedTime"`
	DeletedAt gorm.DeletedAt `gorm: "index"`
}

func (user *User) HashPassword() error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPwd)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err == nil {
		return true
	}
	return false
}
