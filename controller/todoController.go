package controller

//todo控制器层，验证提交的数据。

import (
	"NewTodo/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func TodoIndex(c *gin.Context){

	c.HTML(http.StatusOK, "todo.html", nil)

}

func CreatTodo(c *gin.Context){
	//捞数据
	var todo models.Todo

	c.BindJSON(&todo)
	// fmt.Println(todo)
	// 存入数据库
	// DB.Create(&todo)
	session := sessions.Default(c)

	userInfo := session.Get("userinfo").(string)

	userStruct := models.User{}
	
	json.Unmarshal([]byte(userInfo), &userStruct)

	todo.Uuid = userStruct.Phonenumber
	err := models.CreatATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}


func GetTodoList(c *gin.Context){
	session := sessions.Default(c)

	userInfo := session.Get("userinfo").(string)

	fmt.Println(userInfo)

	userStruct := models.User{}
	
	json.Unmarshal([]byte(userInfo), &userStruct)

	todolist, err := models.GetAllTodo(userStruct.Phonenumber)
	fmt.Println(todolist)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}


func UpdateATodo(c *gin.Context){
	session := sessions.Default(c)

	userInfo := session.Get("userinfo").(string)

	userStruct := models.User{}
	
	json.Unmarshal([]byte(userInfo), &userStruct)

	
	id := c.Params.ByName("id")
	
	todo, err := models.GetATodoByID(id)
	if err != nil {
		return
	}
	
	c.BindJSON(&todo) //把前端发送的request请求中的json数据映射到结构体
	todo.Uuid = userStruct.Phonenumber
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}


func DeleteATodo(c *gin.Context){
	id := c.Params.ByName("id")

	if err := models.DeleteATodoByID(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id": "delete",
		})
	}
}


func LoginOut(c *gin.Context){

	session := sessions.Default(c)

	session.Delete("userinfo")

	session.Save()

	c.Redirect(http.StatusMovedPermanently, "/")

}