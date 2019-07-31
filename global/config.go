package global

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
)

var Config = struct {
	ListenPort  int32
	DbHost      string
	DbPort      int32
	DbUser      string
	DbPwd       string
	DbName      string
	IsSaveLog   bool
	LogPath     string
	Log         *logrus.Logger
	LogLevel    logrus.Level
	VerifyCode  string
	ApiRootPath string
}{
	ListenPort:  8001,
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
	ApiRootPath: fmt.Sprintf("%s/src/github.com/illidan33/wow_api", os.Getenv("GOPATH")),
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
