package routers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"NewTodo/middlewares"
)


func DefaultRouter(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil) //渲染默认的初始化login.html页面

}

func RoutersInit() * gin.Engine{

	r := gin.Default() //利用gin框架生成gin对象

	r.LoadHTMLGlob("templates/*") //加载所有模版
	
	r.Static("../statics", "./statics")  //替换静态文件导入路径

	middlewares.DefaultSessions(r) //cookie实现session的初始化

	r.GET("/", DefaultRouter)

	return r
}

func SetRouters(r * gin.Engine){

	SetLoginRouter(r) //设置login页面相关路由

	SetIndexRouter(r) //设置index页面相关路由

	SetTodoRouter(r)  //设置todo页面相关路由
}