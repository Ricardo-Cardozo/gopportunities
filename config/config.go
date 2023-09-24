package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	return errors.New("fake Error")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
