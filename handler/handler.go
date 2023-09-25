package handler

import (
	"github.com/Ricardo-Cardozo/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializaHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
