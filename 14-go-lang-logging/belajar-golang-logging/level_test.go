package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace") // Tidak keluar di terminal
	logger.Debug("This is Debug") // Tidak keluar di terminal
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // Melakukan set infoLevel mulai dari Trace Level 

	logger.Trace("This is Trace") // Keluar di terminal
	logger.Debug("This is Debug") // Keluar di terminal
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}