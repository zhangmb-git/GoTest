package commonlib

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// var log = logrus.New()

func InitLog() {
	// 设置日志格式为json格式
	//log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)

	//requestLogger := log.WithFields(log.Fields{"request_id": 123, "user_ip": "10.0.72.202"})
	log.Info("something happened on that request")
	log.Warn("something not great happened")

}
