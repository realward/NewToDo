package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"fmt"
	"NewTodo/setting"
)

var (

	DB *gorm.DB

	Token string = "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"

)


func InitMySQL(cfg * setting.MySQLConfig)(err error){

	dsn := fmt.Sprintf(Token, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return err // grom v2

}

func Close(){

	sqlDB, _ := DB.DB()

	sqlDB.Close() //gorm v2
	
}