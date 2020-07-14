package main

import (
	"log"
	"os"

	"github.com/kanapuli/eks-ami-update/config"
)

const (
	configFile = "config.json"
)

func main() {

	// Read and parse config file
	file, err := os.Open(configFile)
	if err != nil {
		log.Printf("Expected opening file %v but got %v", configFile, err)
		os.Exit(1)
	}
	cfg := config.New()
	err = cfg.Read(file)
	if err != nil {
		log.Printf("Expected json parsing for config file %v but got %v", configFile, err)
		os.Exit(1)
	}
}
