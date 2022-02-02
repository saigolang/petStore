package logger

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

// NewLogger setting level to error for the sake of coding exercise
// We can enhance it by sending level as Env variable
func NewLogger() {
	Log = logrus.New()
	Log.SetLevel(logrus.ErrorLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}
