package entity

type User struct {
	Id          int64  `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name"`
	Username    string `json:"username" gorm:"column:username"`
	Age         int    `json:"age" gorm:"column:age"`
	Description string `json:"description" gorm:"column:description"`
}
