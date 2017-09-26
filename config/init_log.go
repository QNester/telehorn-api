package config

import (
	"fmt"
	"os"
	"io"
	"github.com/sirupsen/logrus"
)


// Initialize log configs
func InitLog(mode string) {
	log_path := fmt.Sprintf("../log/%s", Environment())

	if mode == "bot" {
		log_path = fmt.Sprintf("./log/bot_%s.log", Environment())
	} else {
		log_path = fmt.Sprintf("./log/api_%s.log", Environment())
	}
	file, err := os.OpenFile(log_path, os.O_CREATE|os.O_WRONLY, 0666)
	mw := io.MultiWriter(os.Stdout, file)
	if err == nil {
		logrus.Info("Use [%s]", log_path)
		logrus.SetOutput(mw)
	} else {
		logrus.Error(err)
		logrus.Info("Failed to log to file, using default stderr")
	}
	logrus.Info("Logrus loaded loaded successfull")
}