package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.WithField("username", "eko").Info("hello loggin")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "khannedy")
	entry.Info("Hello Entry")
}
