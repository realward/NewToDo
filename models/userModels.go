package models

// user数据库的ORM。

import(
	"NewTodo/dao"
)


func CreatAUser(user *User)(err error){

	err = dao.DB.Create(user).Error

	return

}

func GetAUserByEmail(tag string)(user *User, err error){

	user = new(User)

	if err = dao.DB.Where("email=?", tag).First(user).Error; err != nil {	
		
		return nil, err		
	
	}
	
	return 

}

func GetAUserByPhoneNumber(tag string)(user *User, err error){

	user = new(User)
	
	if err = dao.DB.Where("phonenumber=?", tag).First(user).Error; err != nil {	
		
		return nil, err		
	
	}
	
	return 

}