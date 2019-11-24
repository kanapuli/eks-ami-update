package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kanapuli/eks-ami-update/config"
)

func main() {
	//read config file
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Unable to find config file: %v", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(file)
	cfgFile := new(config.Configuration)
	decoder.Decode(&cfgFile)
	//	fmt.Println(cfgFile)

}
