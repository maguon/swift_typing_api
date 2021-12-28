package common

import (
	"fmt"
	"io"
	"os"
	"swift_typing_api/conf"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetReportCaller(true)
	log.SetLevel(logrus.Level(conf.Config.Log.LogLevel))

	logf, err := rotatelogs.New(
		conf.Config.Log.LogPath+conf.Config.Log.LogName+".%Y-%m-%d",
		rotatelogs.WithLinkName(conf.Config.Log.LogPath+conf.Config.Log.LogName),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		fmt.Println(err)
	}
	log.SetFormatter(&logrus.TextFormatter{DisableColors: false})
	log.SetOutput(io.MultiWriter(os.Stdout, logf))

}

func GetLogger() *logrus.Logger {
	return log
}
