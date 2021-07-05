package modules

import (
	"github.com/illidan33/wow_tools/database"
	"github.com/illidan33/wow_tools/global"
	"time"
)

// 创建登录日志
func createLog(ip string, method string, t uint8) error {
	now := time.Now()
	tm := now.Format("2006-01-02 15:04:05")

	log := database.ApiLoginLog{
		ID:         0,
		IP:         ip,
		Method:     method,
		LoginDate:  now.Format("2006-01-02"),
		Count:      1,
		CreateTime: tm,
		UpdateTime: tm,
		Type:       t,
	}

	err := DbConn.Create(&log).Error
	if err != nil {
		global.Config.Log.Error(err)
	}

	return nil
}

func UpdateOrCreateLog(ip string, method string, t uint8) error {
	date := time.Now().Format("2006-01-02")

	log := database.ApiLoginLog{}
	err := DbConn.Where("ip = ? and method = ? and login_date = ? and type = ?", ip, method, date, t).First(&log).Error
	if err != nil {
		if IsNotFound(err) {
			err = createLog(ip, method, t)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	log.Count += 1

	err = DbConn.Model(&log).Update(database.ApiLoginLog{
		Count:      log.Count,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}).Error
	if err != nil {
		return err
	}

	return nil
}
