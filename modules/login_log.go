package modules

import (
	"github.com/illidan33/wow_api/database"
	"github.com/illidan33/wow_api/global"
	"time"
)

// 创建登录日志
func CreateLog(ip string, method string) error {
	now := time.Now()
	t := now.Format("2006-01-02 15:04:05")

	log := database.LoginLog{
		ID:         0,
		IP:         ip,
		Method:     method,
		LoginDate:  now.Format("2006-01-02"),
		Count:      1,
		CreateTime: t,
		UpdateTime: t,
	}

	err := DbConn.Create(&log).Error
	if err != nil {
		global.Config.Log.Error(err)
	}

	return nil
}

func UpdateOrCreateLog(ip string, method string) error {
	date := time.Now().Format("2006-01-02")

	log := database.LoginLog{}
	err := DbConn.Where("ip = ? and method = ? and login_date = ?", ip, method, date).First(&log).Error
	if err != nil {
		if IsNotFound(err) {
			err = CreateLog(ip, method)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	log.Count += 1

	err = DbConn.Save(&log).Error
	if err != nil {
		return err
	}

	return nil
}
