package utils

import (
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Entry {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logrus.SetReportCaller(true)

	return logrus.WithFields(logrus.Fields{})
}
