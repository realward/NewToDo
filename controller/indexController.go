package controller

//导航控制器层，验证提交的数据。

import (
	"NewTodo/models"
	_"fmt"
	"net/http"
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){

	session := sessions.Default(c)

	userInfo := session.Get("userinfo").(string)

	userStruct := models.User{}
	
	json.Unmarshal([]byte(userInfo), &userStruct)

	c.HTML(http.StatusOK,"index.html",gin.H{
		"name":userStruct.Username,
	})
	// c.HTML(http.StatusOK, "index.html", userInfo)
}

func EnterTodo(c *gin.Context){

	c.Redirect(http.StatusMovedPermanently, "/todo")
}