package global

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
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
	LogLevel  logrus.Level
	// 基础验证
	VerifyCode string
	// 统计天数
	ChartDay int64
}{
	ListenHost: "127.0.0.1",
	ListenPort: 8002,
	DbHost:     "127.0.0.1",
	DbPort:     3306,
	DbUser:     "root",
	DbPwd:      "test123",
	DbName:     "wow_hong",
	IsSaveLog:  true,
	LogPath:    "./logs/log.txt",
	LogLevel:   logrus.DebugLevel,
	VerifyCode: "testcode",
	ChartDay:   20,
}

func init() {
	Log = logrus.New()
	Log.Level = Config.LogLevel
	if Config.IsSaveLog {
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  Config.LogPath,
			logrus.ErrorLevel: Config.LogPath,
			logrus.WarnLevel:  Config.LogPath,
			logrus.DebugLevel: Config.LogPath,
		}
		Log.Hooks.Add(lfshook.NewHook(
			pathMap,
			&logrus.JSONFormatter{},
		))
	}
}
