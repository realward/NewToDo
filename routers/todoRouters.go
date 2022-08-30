package routers

// todo 页面的路由设计

import (
	"NewTodo/controller"
	"NewTodo/middlewares"

	"github.com/gin-gonic/gin"
)

func SetTodoRouter(r *gin.Engine) {

	r.GET("/todo", controller.TodoIndex)

	TodoRouterGroup := r.Group("v1") //前端设置，现在这里无法更改

	TodoRouterGroup.Use(middlewares.AuthSession())
	{
		
		TodoRouterGroup.POST("/todo", controller.CreatTodo) //添加

		TodoRouterGroup.GET("/todo", controller.GetTodoList) //查看全部

		TodoRouterGroup.PUT("/todo/:id", controller.UpdateATodo) //修改某一个

		TodoRouterGroup.DELETE("/todo/:id", controller.DeleteATodo) //删除某一个

	}

	r.POST("/todo", controller.LoginOut) //退出

}
