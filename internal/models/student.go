package models

type Student struct {
ID        int    `json :"id" gorm:"primarykey"`
FirstName string `json : "firstname"`
LastName  string `json : "lastname"`
Age       int    `json: "age"`
}
