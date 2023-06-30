package log

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	// Set the Logrus formatter
	Logger.SetFormatter(&logrus.JSONFormatter{})

	// Set the Logrus level
	Logger.SetLevel(logrus.InfoLevel)
}
