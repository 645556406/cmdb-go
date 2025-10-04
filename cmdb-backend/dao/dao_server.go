package dao

import (
	"cmdb-backend/model"
	"fmt"
	"log"
)

type ServerStatusNum struct {
	Total   int `json:"total"`
	Online  int `json:"online"`
	Offline int `json:"offline"`
}

func GetServerList() []model.Server {

	var serverList []model.Server
	err := DB.Where("deleted_at IS NULL").Find(&serverList).Error
	if err != nil {
		log.Println(err)
	}
	return serverList
}

func UpdateServer(server model.Server) error {
	// 使用 map 的简洁写法（推荐）
	updates := make(map[string]interface{})

	// 一行一个字段，清晰简洁
	if server.Hostname != "" {
		updates["hostname"] = server.Hostname
	}
	if server.CPU > 0 {
		updates["cpu"] = server.CPU
	}
	if server.Memory > 0 {
		updates["memory"] = server.Memory
	}
	if server.IP != "" {
		updates["ip"] = server.IP
	}
	if server.Username != "" {
		updates["username"] = server.Username
	}
	if server.Password != "" {
		updates["password"] = server.Password
	}
	if server.Area != "" {
		updates["area"] = server.Area
	}
	if server.PublicKey != "" {
		updates["public_key"] = server.PublicKey
	}
	if server.Env != "" {
		updates["env"] = server.Env
	}
	if server.OS != "" {
		updates["os"] = server.OS
	}
	if server.Owner != "" {
		updates["owner"] = server.Owner
	}
	if server.Port > 0 {
		updates["port"] = server.Port
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	result := DB.Model(&model.Server{}).Where("id = ?", server.ID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("server not found")
	}

	return nil
}

func AddServer(server model.Server) error {
	err := DB.Create(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func DelServer(server model.Server) error {
	err := DB.Delete(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetServerDetailByID(id uint64) (model.Server, error) {
	var server model.Server
	err := DB.First(&server, id).Error
	if err != nil {
		log.Println(err)
		return server, err
	} else {
		return server, nil
	}
}

func GetServerOneByIP(ip string) (model.Server, error) {

	var server model.Server
	err := DB.Where("IP = ?", ip).First(&server).Error
	if err != nil {
		log.Println(err)
		return server, err
	} else {
		return server, nil
	}
}

func GetServerCount() (ServerStatusNum, error) {
	var servers []model.Server
	var serverStatusNum ServerStatusNum
	result := DB.Find(&servers).RowsAffected
	resultOnline := DB.Where("status=?", 1).Find(&servers).RowsAffected
	resultOffline := DB.Where("status=?", 0).Find(&servers).RowsAffected
	serverStatusNum.Total = int(result)
	serverStatusNum.Online = int(resultOnline)
	serverStatusNum.Offline = int(resultOffline)
	return serverStatusNum, nil
}

func GetOnlineCountServer() (int, error) {
	var servers []model.Server
	result := DB.Where("status=?", 1).Find(&servers).RowsAffected
	return int(result), nil
}

func GetOfflineCountServer() (int, error) {
	var servers []model.Server
	result := DB.Where("status=?", 0).Find(&servers).RowsAffected
	return int(result), nil
}

func GetServerIPList() ([]model.Server, error) {
	var servers []model.Server
	DB.Select("ID", "IP").Find(&servers)
	return servers, nil
}

func UpdateServerStatus(id uint, s int) {
	var servers model.Server
	servers.Status = s
	err := DB.Model(&model.Server{}).Where("id = ?", id).Select("Status").Updates(servers).Error
	if err != nil {
		log.Println(err)
	}
	return
}
