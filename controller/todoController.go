package controller

//todo控制器层，验证提交的数据。

import(
	"NewTodo/models"
	"net/http"
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
	todolist, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}


func UpdateATodo(c *gin.Context){
	id := c.Params.ByName("id")
	
	todo, err := models.GetATodoByID(id)
	if err != nil {
		return
	}

	c.BindJSON(&todo) //把前端发送的request请求中的json数据映射到结构体

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
