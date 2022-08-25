package controller

//登录控制器层，验证提交的数据。
import (
	"NewTodo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRegister(c *gin.Context) {

	tag := c.PostForm("tag")

	if tag == "login" {

		Login(c)

	} else if tag == "register" {

		Register(c)

	}
}

func Login(c *gin.Context) {
	// 1.数据库查询该账户是否已经注册
	LoginName := c.PostForm("loginUsername")

	LoginPassword := c.PostForm("loginPassword")

	var userInfo *models.User

	var err error

	for {

			if userInfo, err = models.GetAUserByEmail(LoginName); nil == err {

				break

			}else if userInfo, err = models.GetAUserByPhoneNumber(LoginName); nil == err {

				break

			}else{
				//账户不存在
				c.JSON(http.StatusUnauthorized, gin.H{
					"status": "账户不存在",
				})

				c.Redirect(http.StatusMovedPermanently, "/")

				return
			}
		}
	
	if userInfo.Password == LoginPassword {

		c.Set("userinfo", userInfo)
		fmt.Println(userInfo)
		c.Next() //执行新建session中间件

		c.Redirect(http.StatusMovedPermanently, "/index")//登录
	}else{

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "密码错误",
		})

		c.Redirect(http.StatusMovedPermanently, "/")

		return
	}


}



func Register(c *gin.Context) {

	u := &models.User{}

	c.ShouldBind(u)

	// 1.注册的信息是否合法

	// 2.数据库查询该账户是否已经注册

	// 3.回到登陆页面

}
