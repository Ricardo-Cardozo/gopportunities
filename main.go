package main

import (
	"github.com/Ricardo-Cardozo/gopportunities/config"
	"github.com/Ricardo-Cardozo/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("INITIALIZATION_ERROR: %v", err)
		return
	}

	router.Initialize()
}
