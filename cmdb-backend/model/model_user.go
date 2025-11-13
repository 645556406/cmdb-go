package model

import "gorm.io/gorm"

// User 用户表结构
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
