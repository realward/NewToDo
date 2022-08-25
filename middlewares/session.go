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

	r.Use(sessions.Sessions("sessionId", store))

}//注册全局session

func SetSessions(c *gin.Context) {
	fmt.Println("SetSessions start")
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

	userinfo, _ :=json.Marshal(userInfo)

	session.Set("userinfo",string(userinfo))

	err := session.Save()

	if err == nil{
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "session successfully",
		// })
	}else {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "session failed",
		// })
	}
}