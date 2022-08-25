package controller

//登录控制器层，验证提交的数据。
import (
	"NewTodo/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginRegister(c *gin.Context) {

	tag := c.PostForm("tag") //通过判断发送的post请求中，type为hidden的<input>标签中，tag项对应的值，来区分post请求是login还是register

	if tag == "login" {

		Login(c)

	} else if tag == "register" {

		Register(c)

	}
}

func Login(c *gin.Context) {
	// 1.数据库查询该账户是否已经注册
	LoginName := c.PostForm("loginUsername") //前端提交的表单中name=loginUsername所对应的填写值

	LoginPassword := c.PostForm("loginPassword") //前端提交的表单中name=loginPassword所对应的填写值

	var userInfo *models.User

	var err error

	for {

		if userInfo, err = models.GetAUserByEmail(LoginName); nil == err {

			break

		} else if userInfo, err = models.GetAUserByPhoneNumber(LoginName); nil == err {

			break

		} else {
			//账户不存在
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "账户不存在",
			})

			c.Redirect(http.StatusMovedPermanently, "/")

			return
		}
	}

	if userInfo.Password == LoginPassword {

		c.Set("userinfo", userInfo) //将前端获取的用户名和密码注入gin.context中

		c.Next() //执行新建session的中间件

		c.Redirect(http.StatusMovedPermanently, "/index") //完成登录，进入index页面

	} else {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "密码错误",
		})

		c.Redirect(http.StatusMovedPermanently, "/") //重定向到起始login页面

		return
	}

}

func Register(c *gin.Context) {

	registerEmail := c.PostForm("registerEmail")

	registerPhoneNumber := c.PostForm("registerPhoneNumber")
	// 1.注册的信息是否合法

	if !models.FilterEmail(registerEmail){
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "邮箱不合法",
		})		
		c.Redirect(http.StatusMovedPermanently, "/") 
		return
	}
	if !models.FilterPhonenumber(registerPhoneNumber){
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "电话号码不合法",
		})		
		c.Redirect(http.StatusMovedPermanently, "/") 
		return
	}
	
	// 2.数据库查询该账户是否已经注册
	registerPassword := models.Md5(c.PostForm("registerPassword"))
	
	user := &models.User{
		Username : c.PostForm("registerName"),
		Password : registerPassword,
		Email	 : registerEmail,
		Phonenumber: registerPhoneNumber,
		Status: true,
	}

	var userInfo *models.User

	var err error

	if userInfo, err = models.GetAUserByRegisterInfo(user); err == nil{

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "用户存在",
		})

		c.Redirect(http.StatusMovedPermanently, "/") 
		return
	}

	if err = models.CreatAUser(userInfo); err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "用户创建失败",
		})
		c.Redirect(http.StatusMovedPermanently, "/") 
		return
	}
	// 3.回到登陆页面

	c.Redirect(http.StatusMovedPermanently, "/") 

}
