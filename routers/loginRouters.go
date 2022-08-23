package routers

// 登录页面的路由设计
import(
	"github.com/gin-gonic/gin"
	"NewTodo/controller"
)

func SetLoginRouter(r *gin.Engine ){

	LoginRouterGroup := r.Group("/login")
	{
		LoginRouterGroup.GET("/", DefaultRouter)

		LoginRouterGroup.POST("/", controller.LoginRegister)

	}

}

