package model

import "gorm.io/gorm"

// 部门表
type Departments struct {
	gorm.Model
	Name        string `gorm:"column:name"`         // 部门名称
	Desc        string `gorm:"column:desc"`         // 部门描述
	ParentId    uint   `gorm:"column:parent_id"`    // 父部门ID
	LeaderId    uint   `gorm:"column:leader_id"`    // 负责人ID
	LeaderName  string `gorm:"column:leader_name"`  // 负责人名称
	LeaderEmail string `gorm:"column:leader_email"` // 负责人邮箱
	LeaderPhone string `gorm:"column:leader_phone"` // 负责人电话
}
