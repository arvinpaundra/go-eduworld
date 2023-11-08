package utils

import (
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-02-02 15:04:05",
		PrettyPrint:     true,
	})
	logrus.SetReportCaller(true)

	return logrus.WithFields(logrus.Fields{})
}
