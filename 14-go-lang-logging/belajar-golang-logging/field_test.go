package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// # Menambahkan field tambahan secara manual dan satu persatu
func TestField(t *testing.T) {
	logger := logrus.New()

	// # Mengubah log menjadi format json
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "khannedy").Info("Hello World")
	logger.WithField("username", "eko").WithField("name", "eko kurniawan").Info("Hello World")
}

// # Menambahkan field tambahan secara banyak dan tidak satu persatu
func TestFields(t *testing.T) {
	logger := logrus.New()

	// # Mengubah log menjadi format json
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "eko",
		"address":  "Eko Kurniawan",
	}).Info("Hello World")
}
