package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)
func DefaultSessions(r * gin.Engine){

	store := cookie.NewStore([]byte("secret")) //cookie实现，后面改成redis

	r.Use(sessions.Sessions("mysession", store))

}//注册全局session

func SetSessions() gin.HandlerFunc  {

	return func(c  *gin.Context){
		
		userInfo, ok := c.Get("userinfo")

		fmt.Println(userInfo)
		if !ok{
			//日志记录
			c.JSON(http.StatusOK, gin.H{
				"status": "获取用户信息失败",
			})
			return
		}
		session := sessions.Default(c)

		userinfo, _ := json.Marshal(userInfo) //结构体编码为二进制的json数据流

		session.Set("userinfo",string(userinfo)) //session保存的k-v值中的v只能是字符串类型

		err := session.Save() //必须通过save来使set生效

		if err != nil{
			c.JSON(http.StatusOK, gin.H{
				"status": "session failed",
			})
		}
	}
}

func AuthSession() gin.HandlerFunc{
	return func(c *gin.Context) {

		session := sessions.Default(c)
	
		if session.Get("userinfo") == nil{
	
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "账户信息session不存在",
			})
	
			c.Redirect(http.StatusMovedPermanently, "/")

			c.Abort()
	
			return
		}

		c.Next()
	}
}