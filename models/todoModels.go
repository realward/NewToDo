package models

// todolist数据库的ORM。


import(
	"NewTodo/dao"
)



func CreatATodo(todo *Todo)(err error){
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo()(todolist []*Todo, err error){

	if err = dao.DB.Find(&todolist).Error; err!= nil{
		return nil, err
	}
	return todolist, err
}
	
func GetATodoByID(id string)(todo *Todo, err error){

	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err		
		// c.JSON(http.StatusOK, gin.H{
		// 	"error": err.Error(),
		// })
	}
	return 
}

func UpdateATodo(todo *Todo)(err error){
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodoByID(id string)(err error){
	err = dao.DB.Where("id=?",id).Delete(&Todo{}).Error
	return
}