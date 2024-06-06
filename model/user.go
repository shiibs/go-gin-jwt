package model

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"column:email;unique"`
	Password string `json:"_" gorm:"column:password"`
}