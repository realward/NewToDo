package controller

//登录控制器层，验证提交的数据。
import (
	"NewTodo/models"

	"github.com/gin-gonic/gin"
)
func LoginRegister(c *gin.Context){

	tag := c.PostForm("tag")

	if tag == "login" {

		Login(c)

	}else if tag == "register" {

		Register(c)
	}
}

func Login(c *gin.Context) {
	// 1.数据库查询该账户是否已经注册
}

func Register(c *gin.Context) {

	u := &models.User{}

	c.ShouldBind(u)

}