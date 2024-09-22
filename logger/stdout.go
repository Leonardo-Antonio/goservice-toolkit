package logger

import (
	"io"

	logOrigin "log"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func New(out io.Writer) {
	log = logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		PrettyPrint: true,
		DataKey:     "data",
	}
	log.Out = out
	logOrigin.Println("✅ Log configured")
}

type (
	level string
	Input struct {
		Message string         `json:"message"`
		Data    map[string]any `json:"data"`
		Level   level          `json:"type"`
	}
)

const (
	Debug level = "debug"
	Info  level = "info"
	Warn  level = "warn"
	Error level = "error"
)

func Write(in Input) {
	var level logrus.Level
	switch in.Level {
	case Debug:
		level = logrus.DebugLevel
	case Info:
		level = logrus.InfoLevel
	case Warn:
		level = logrus.WarnLevel
	case Error:
		level = logrus.ErrorLevel
	}

	if in.Data == nil {
		log.WithFields(nil).Log(level, in.Message)
		return
	}

	log.WithFields(in.Data).Log(level, in.Message)
}
