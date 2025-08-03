package dao

import (
	"cmdb-backend/model"
	"database/sql"
	"log"
)

type ServerStatusNum struct {
	Total   int `json:"total"`
	Online  int `json:"online"`
	Offline int `json:"offline"`
}

func GetServerList() []model.Server {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var serverList []model.Server
	err := db.Where("deleted_at IS NULL").Find(&serverList).Error
	if err != nil {
		log.Println(err)
	}
	return serverList
}

func UpdateServer(server model.Server) error {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	serverUpdate := model.Server{}
	serverFields := make([]string, 0)
	if server.ID >= 0 {
		serverUpdate.ID = server.ID
		serverFields = append(serverFields, "id")
	}
	if server.Hostname != "" {
		serverUpdate.Hostname = server.Hostname
		serverFields = append(serverFields, "HostName")
	}
	if server.CPU > 0 {
		serverUpdate.CPU = server.CPU
		serverFields = append(serverFields, "CPU")
	}
	if server.Memory > 0 {
		serverUpdate.Memory = server.Memory
		serverFields = append(serverFields, "Memory")
	}
	if server.IP != "" {
		serverUpdate.IP = server.IP
		serverFields = append(serverFields, "IP")
	}
	if server.IP != "" {
		serverUpdate.Username = server.Username
		serverFields = append(serverFields, "Username")
	}
	if server.IP != "" {
		serverUpdate.Password = server.Password
		serverFields = append(serverFields, "Password")
	}
	if server.IP != "" {
		serverUpdate.Area = server.Area
		serverFields = append(serverFields, "Area")
	}
	if server.IP != "" {
		serverUpdate.PublicKey = server.PublicKey
		serverFields = append(serverFields, "PublicKey")
	}
	if server.Env != "" {
		serverUpdate.Env = server.Env
		serverFields = append(serverFields, "Env")
	}
	if server.OS != "" {
		serverUpdate.OS = server.OS
		serverFields = append(serverFields, "OS")
	}
	if server.Owner != "" {
		serverUpdate.Owner = server.Owner
		serverFields = append(serverFields, "Owner")
	}
	if server.Port > 0 {
		serverUpdate.Port = server.Port
		serverFields = append(serverFields, "Port")
	}
	err := db.Model(&model.Server{}).Where("id = ?", server.ID).Select(serverFields).Updates(serverUpdate).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func AddServer(server model.Server) error {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	err := db.Create(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func DelServer(server model.Server) error {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	err := db.Delete(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetServerDetailByID(id uint64) (model.Server, error) {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var server model.Server
	err := db.First(&server, id).Error
	if err != nil {
		log.Println(err)
		return server, err
	} else {
		return server, nil
	}
}

func GetServerOneByIP(ip string) (model.Server, error) {
	db := NewDB()
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var server model.Server
	err := db.Where("IP = ?", ip).First(&server).Error
	if err != nil {
		log.Println(err)
		return server, err
	} else {
		return server, nil
	}
}

func GetServerCount() (ServerStatusNum, error) {
	db := NewDB()
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var servers []model.Server
	var serverStatusNum ServerStatusNum
	result := db.Find(&servers).RowsAffected
	resultOnline := db.Where("status=?", 1).Find(&servers).RowsAffected
	resultOffline := db.Where("status=?", 0).Find(&servers).RowsAffected
	serverStatusNum.Total = int(result)
	serverStatusNum.Online = int(resultOnline)
	serverStatusNum.Offline = int(resultOffline)
	return serverStatusNum, nil
}

func GetOnlineCountServer() (int, error) {
	db := NewDB()
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var servers []model.Server
	result := db.Where("status=?", 1).Find(&servers).RowsAffected
	return int(result), nil
}

func GetOfflineCountServer() (int, error) {
	db := NewDB()
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {

		}
	}(sqlDB)
	var servers []model.Server
	result := db.Where("status=?", 0).Find(&servers).RowsAffected
	return int(result), nil
}

func GetServerIPList() ([]model.Server, error) {
	db := NewDB()
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println(err)
		}
	}(sqlDB)
	var servers []model.Server
	db.Select("ID", "IP").Find(&servers)
	return servers, nil
}

func UpdateServerStatus(id uint, s int) {
	db := NewDB()
	sqlDB, errDb := db.DB()
	if errDb != nil {
		log.Println(errDb)
	}
	// Close
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println(err)
		}
	}(sqlDB)
	var servers model.Server
	servers.Status = s
	err := db.Model(&model.Server{}).Where("id = ?", id).Select("Status").Updates(servers).Error
	if err != nil {
		log.Println(err)
	}
	return
}
