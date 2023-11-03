package main

import (
	"log"

	"llm_plateform_resourcemanagement/config"
	"llm_plateform_resourcemanagement/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
