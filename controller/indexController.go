package controller

//导航控制器层，验证提交的数据。

import (
	"NewTodo/models"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){

	session := sessions.Default(c)

	var userInfo string

	if session.Get("userinfo") == nil{

		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "账户信息不存在",
		})

		c.Redirect(http.StatusMovedPermanently, "/")

		return
	}
	userInfo = session.Get("userinfo").(string)
	fmt.Println(userInfo)
	userStruct := models.User{}
	json.Unmarshal([]byte(userInfo), &userStruct)
	fmt.Println(userStruct)
	c.HTML(http.StatusOK,"index.html",gin.H{
		"name":userStruct.Username,
	})
	// c.HTML(http.StatusOK, "index.html", userInfo)
}

func EnterTodo(c *gin.Context){


}