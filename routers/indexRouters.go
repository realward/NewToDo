package routers

//导航页面的路由设计
import(
	"github.com/gin-gonic/gin"
	"NewTodo/controller"
)


func SetIndexRouter(r *gin.Engine) {

	IndexRouterGrout := r.Group("/index")
	{
		IndexRouterGrout.GET("/", controller.Index)

		IndexRouterGrout.POST("/", controller.EnterTodo)
	}

}
