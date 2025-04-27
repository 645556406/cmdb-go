package model

import "gorm.io/gorm"

type Server struct {
	gorm.Model
	Hostname  string `gorm:"unique;not null"` // 主机名
	IP        string `gorm:"unique"`          // IP 地址
	Port      int    // 端口号
	Password  string // 密码
	PublicKey string // 公钥
	OS        string // 操作系统
	CPU       int    // CPU 核数
	Memory    int    // 内存（GB）
	Area      string // 地区
	Owner     string // 负责人
	Env       string // 环境（prod/dev/test）
	Status    int    // 状态（0: 未使用, 1: 使用中）
}
