package dao

import (
	"cmdb-backend/model"
	"log"
)

func GetServerList() []model.Server {
	db := NewDB()
	var serverList []model.Server
	err := db.Find(&serverList).Where("delete_at = ", "").Error
	if err != nil {
		log.Println(err)
	}
	return serverList
}

func UpdateServer(server model.Server) error {
	db := NewDB()
	serverUpdate := model.Server{}
	serverFields := make([]string, 0)
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
		serverFields = append(serverFields, "CPU")
	}
	if server.IP != "" {
		serverUpdate.IP = server.IP
		serverFields = append(serverFields, "IP")
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
	err := db.Model(&model.Server{}).Where("id = ?", server.ID).Select(serverFields).Updates(serverUpdate).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func AddServer(server model.Server) error {
	db := NewDB()
	err := db.Create(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func DelServer(server model.Server) error {
	db := NewDB()
	err := db.Delete(&server).Error
	if err != nil {
		log.Println(err)
	}
	return err
}

func GetServerDetailByID(id uint64) (model.Server, error) {
	db := NewDB()
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
	var server model.Server
	err := db.Where("IP = ?", ip).First(&server).Error
	if err != nil {
		log.Println(err)
		return server, err
	} else {
		return server, nil
	}
}
