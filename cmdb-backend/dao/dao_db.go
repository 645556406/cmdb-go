package dao

import (
	"cmdb-backend/model"
	"cmdb-backend/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	config := utils.LoadYamlConfigNew("config/dev.yaml")
	m := config["DB"].(map[string]interface{})
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m["username"], m["password"], m["hostname"], m["port"], m["database"])
	db, errDb := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDb != nil {
		panic(errDb)
	}
	errAutoMigrate := db.AutoMigrate(&model.Server{})
	if errAutoMigrate != nil {
		log.Println(errAutoMigrate)
	}
	return db
}
