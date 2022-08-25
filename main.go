package main

import (
	"NewTodo/dao"
	"NewTodo/models"
	"NewTodo/routers"
	"NewTodo/setting"
	"fmt"
	_ "log"

	_ "github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	_ "github.com/gin-gonic/gin"
)



func main(){

	file := "./config/config.ini"

	if err := setting.SettingInit(file); err != nil {

		fmt.Printf("config init failed, err : %v\n", err)

		// 记录日志

	}

	if err := dao.InitMySQL(setting.Conf.MySQLConfig); err != nil {
		
		fmt.Printf("mysql init failed, err : %v\n", err)
		
		// 记录日志

		return
	}

	defer dao.Close()

	// 模型绑定dao.DB.AutoMigrate(&models.Todo{})
	dao.DB.AutoMigrate(&models.User{},&models.Todo{},)
	
	r := routers.RoutersInit()

	routers.SetRouters(r)

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {

		fmt.Printf("server startup failed, err:%v\n", err)

		//记录日志
	}
}

