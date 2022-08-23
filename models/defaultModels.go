package models

type User struct {
	Username string 	`json:"registerName" form:"registerName"`
	Password string 	`json:"registerPassword" form:"registerPassword"`
	Email string 		`json:"registerEmail" form:"registerEmail"`
	Phonenumber string 	`json:"registerPhoneNumber" form:"registerPhoneNumber" gorm:"primaryKey"`
	Todos []Todo 		`gorm:"ForeignKey:Fid;AssociationForeignKey:Phonenumber"`
}


type Todo struct {
	Id     int    `json:"id" form:"id"`
	Title  string `json:"title" form:"title"`
	Status bool   `json:"status" form:"status"`
	Phonenumber string `json:"phonenumber" form:"phonenumber"`
}

func (u User) TableName() string {

	return "Users"

}

func (t Todo) TableName() string {

	return "Todos"

}