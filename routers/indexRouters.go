package routers

//导航页面的路由设计
import(
	"NewTodo/controller"
	"NewTodo/middlewares"

	"github.com/gin-gonic/gin"
)


func SetIndexRouter(r *gin.Engine) {

	IndexRouterGroup := r.Group("/index")

	IndexRouterGroup.Use(middlewares.AuthSession())
	{
		IndexRouterGroup.GET("/", controller.Index)

		IndexRouterGroup.POST("/", controller.EnterTodo)
	}

}
