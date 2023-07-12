package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func GetLogger() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)

	return log
}
