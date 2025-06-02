package model

import "gorm.io/gorm"

// Role 角色表结构
type Role struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}
