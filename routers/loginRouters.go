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
		LoginRouterGroup.GET("/", DefaultRouter) // 设置/login/的GET请求路由

		LoginRouterGroup.POST("/", controller.LoginRegister, middlewares.SetSessions())//设置/login/的post请求路由

	}

}

