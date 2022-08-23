package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
)


func DefaultRouter(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

}

func RoutersInit() * gin.Engine{

	r := gin.Default() //利用gin框架生成gin对象

	r.LoadHTMLGlob("templates/*") //加载所有模版
	
	r.Static("../statics", "./statics")  //替换静态文件导入路径

	r.GET("/", DefaultRouter)

	return r
}

func SetRouters(r * gin.Engine){

	SetLoginRouter(r)

	SetIndexRouter(r)

	SetTodoRouter(r)
}