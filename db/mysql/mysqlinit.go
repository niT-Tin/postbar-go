package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"postbar/datamodels"
	"postbar/db"
	err2 "postbar/err"
	"postbar/utils"
)

func GetMySQLDB() *gorm.DB {
	return db.MysqlDB
}

func init() {
	var err error
	env := utils.GetMySqlENV()
	db.MysqlUrl = env.User + ":" + env.Password + "@tcp(" + env.Host + ":" + env.Port + ")/" + env.DBName + "charset=utf8mb4&parseTime=True&loc=Local"
	db.MysqlDB, err = gorm.Open(mysql.Open(db.MysqlUrl), &gorm.Config{})
	if err2.Reciteerr(&err) {
		return
	}
	err = db.MysqlDB.AutoMigrate(&datamodels.Error{})
	if err2.Reciteerr(&err) {
		return
	}
	err = db.MysqlDB.AutoMigrate(&datamodels.Router{})
	if err2.Reciteerr(&err) {
		return
	}
}
