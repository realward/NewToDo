package routers

// 登录页面的路由设计
import (
	"NewTodo/controller"
	"NewTodo/middlewares"

	"github.com/gin-gonic/gin"
)

func SetLoginRouter(r *gin.Engine ){

	LoginRouterGroup := r.Group("/login")
	{
		LoginRouterGroup.GET("/", DefaultRouter)

		LoginRouterGroup.POST("/", controller.LoginRegister, middlewares.SetSessions)

	}

}

