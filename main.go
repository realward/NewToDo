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

	if err := setting.SettingInit(file); err != nil {  //载入配置文件

		fmt.Printf("config init failed, err : %v\n", err)

		// 记录日志

	}

	if err := dao.InitMySQL(setting.Conf.MySQLConfig); err != nil { //初始化数据库
		
		fmt.Printf("mysql init failed, err : %v\n", err)
		
		// 记录日志

		return
	}

	defer dao.Close()

	// 模型绑定dao.DB.AutoMigrate(&models.Todo{})
	dao.DB.AutoMigrate(&models.User{},&models.Todo{},)//数据结构映射到数据库
	
	r := routers.RoutersInit()//初始化路由

	routers.SetRouters(r) //设置路由 

	if err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {

		fmt.Printf("server startup failed, err:%v\n", err)

		//记录日志
	}
}

