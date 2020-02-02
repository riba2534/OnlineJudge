package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/riba2534/OnlineJudge/backend/pkg/util"
)

type User struct {
	gorm.Model
	UserName  string    `json:"username" gorm:"unique"`
	PassWord  string    `json:"password"`
	Name      string    `json:"name"`
	RegTime   time.Time `json:"regtime"`
	LoginTime time.Time `json:"logintime"`
	School    string    `json:"school"`
	Course    string    `json:"course"`
	Classes   string    `json:"classes"`
	Number    string    `json:"number"`
	RealName  string    `json:"realname"`
	QQ        string    `json:"qq"`
	Email     string    `json:"email"`
	Type      int       `json:"type"` // 1 普通, 2 管理员, 3 超级管理员
}

type UserData struct {
	gorm.Model
	UserName string `json:"username" gorm:"unique"`
	AC       int    `json:"ac"`
	Submit   int    `json:"submit"`
	Score    int    `json:"score"`
	Des      string `json:"des"`
	Rating   int    `json:"rating"`
}

// 获取user_data信息
func GetUserDataByUserName(username string) ([]UserData, error) {
	var userData []UserData
	err := db.Where(UserData{UserName: username}).Find(&userData).Error
	if err != nil {
		return nil, err
	}
	return userData, nil
}

// 用于登陆
func CheckAuth(username, password string) (bool, error) {
	// 密码为md5加密形式
	password = util.EncodeMD5(password)
	var user User
	err := db.Select("id").Where(User{UserName: username, PassWord: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 检测用户是否存在
func CheckUser(username string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{UserName: username}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 创建用户
func CreateUser(user User) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(&UserData{
		UserName: user.UserName,
		AC:       0,
		Submit:   0,
		Score:    0,
		Des:      "",
		Rating:   0,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
