package global

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Config = struct {
	ListenPort int32
	DbHost     string
	DbPort     int32
	DbUser     string
	DbPwd      string
	DbName     string
	IsSaveLog  bool
	LogPath    string
	Log        *logrus.Logger
	LogLevel   logrus.Level
	VerifyCode string
}{
	ListenPort: 8001,
	DbHost:     "127.0.0.1",
	DbPort:     3306,
	DbUser:     "test",
	DbPwd:      "test",
	DbName:     "wow_hong",
	IsSaveLog:  false,
	Log:        logrus.New(),
	LogPath:    "./logs/log.txt",
	LogLevel:   logrus.DebugLevel,
	VerifyCode: "testcode",
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
