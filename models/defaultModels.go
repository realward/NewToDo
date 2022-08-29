package models

type User struct {
	Username string 	`json:"registerName" form:"registerName" gorm:"column:username"`
	Password string 	`json:"registerPassword" form:"registerPassword" gorm:"column:password"`
	Email string 		`json:"registerEmail" form:"registerEmail" gorm:"column:email"`
	Phonenumber string 	`json:"registerPhoneNumber" form:"registerPhoneNumber" gorm:"primaryKey;column:phonenumber"`
	Status bool  	   	`json:"status" form:"status" gorm:"column:status"`
	Todos []Todo 		`gorm:"ForeignKey:Uuid;AssociationForeignKey:Phonenumber;"`
}


type Todo struct {
	Id     int    `json:"id" form:"id" gorm:"column:id"`
	Title  string `json:"title" form:"title" gorm:"column:title"`
	Status bool   `json:"status" form:"status" gorm:"column:status"`
	Uuid string   `json:"Uuid" form:"Uuid" gorm:"column:Uuid;size:255"`
}

func (u User) TableName() string {

	return "users"

}

func (t Todo) TableName() string {

	return "todos"

}