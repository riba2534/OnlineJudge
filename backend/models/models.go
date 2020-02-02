package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/riba2534/OnlineJudge/backend/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error

	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	// 全局禁用复数表名
	db.SingularTable(true)
	// 最大链接数
	db.DB().SetMaxIdleConns(10)
	// 最大打开数
	db.DB().SetMaxOpenConns(100)
	// 自动迁移
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4").AutoMigrate(
		&User{},
		&UserData{},
	)
}

func CloseDB() {
	defer db.Close()
}
