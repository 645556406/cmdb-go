package dao

import (
	"cmdb-backend/model"
	"cmdb-backend/utils"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	config := utils.LoadYamlConfigNew("config/dev.yaml")
	m := config["DB"].(map[string]interface{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m["username"], m["password"], m["hostname"], m["port"], m["database"])
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic(errDb)
	}
	errAutoMigrate := db.AutoMigrate(&model.Server{}, &model.Role{}, &model.Departments{}, &model.User{})
	if errAutoMigrate != nil {
		log.Println(errAutoMigrate)
	}
	// 获取底层的sql.DB对象
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	return db
}

type Database struct {
	db *gorm.DB
}

func InitDB() (*Database, error) {
	config := utils.LoadYamlConfigNew("config/dev.yaml")
	m := config["DB"].(map[string]interface{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m["username"], m["password"], m["hostname"], m["port"], m["database"])
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic(errDb)
	}
	errAutoMigrate := db.AutoMigrate(&model.Server{}, &model.Role{}, &model.Departments{}, &model.User{})
	if errAutoMigrate != nil {
		log.Println(errAutoMigrate)
	}
	// 获取底层的sql.DB对象
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	return &Database{db: db}, nil
}
func (d *Database) GetDB() *gorm.DB {
	return d.db
}
