package global

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Config = struct {
	// 系统
	ListenHost  string
	ListenPort  int32
	ApiRootPath string
	// 数据库
	DbHost string
	DbPort int32
	DbUser string
	DbPwd  string
	DbName string
	// 日志
	IsSaveLog bool
	LogPath   string
	Log       *logrus.Logger
	LogLevel  logrus.Level
	// 简单验证
	VerifyCode string
	// 统计天数
	ChartDay int64
}{
	ListenHost:  "127.0.0.1",
	ListenPort:  8002,
	ApiRootPath: "/data/golang/go/src/github.com/illidan33/wow_tools",
	DbHost:      "127.0.0.1",
	DbPort:      3306,
	DbUser:      "testU",
	DbPwd:       "testP",
	DbName:      "wow_hong",
	IsSaveLog:   false,
	Log:         logrus.New(),
	LogPath:     "./logs/log.txt",
	LogLevel:    logrus.DebugLevel,
	VerifyCode:  "testcode",
	ChartDay:    20,
}

func init() {
	if Config.IsSaveLog {
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  Config.LogPath,
			logrus.ErrorLevel: Config.LogPath,
			logrus.WarnLevel:  Config.LogPath,
		}
		Config.Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))
	}
	Config.Log.Level = Config.LogLevel
}
