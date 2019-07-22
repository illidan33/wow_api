package global

import "github.com/sirupsen/logrus"

var Config = struct {
	ListenPort int32
	DbHost     string
	DbPort     int32
	DbUser     string
	DbPwd      string
	DbName     string
	LogPath   string
	Log    *logrus.Logger

}{
	ListenPort: 8001,
	DbHost:     "127.0.0.1",
	DbPort:     3306,
	DbUser:     "test",
	DbPwd:      "test",
	DbName:     "wow_hong",
	LogPath:   "./logs/log.txt",
}
