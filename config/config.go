package config

import (
	"encoding/json"
	"io"
)

// Configuration is the base application configuration
type Configuration struct {
	AutoscalingEnabled  bool   `json:"autoscalingEnabled"`
	AutoscalerNamespace string `json:"autoscalerNamespace"`
	AsgName             string `json:"asgName"`
	AsgDesiredCapacity  int    `json:"asgDesiredCapacity"`
	DryRun              bool   `json:"dryRun"`
}

// New returns a new instance of Configuration
func New() *Configuration {
	return &Configuration{
		AutoscalingEnabled:  false,
		AutoscalerNamespace: "default",
		AsgName:             "",
		AsgDesiredCapacity:  0,
		DryRun:              true,
	}
}

// Read reads config file and populates Configuration struct
func (c *Configuration) Read(file io.Reader) error {
	json.NewDecoder(file).Decode(&c)
	return nil
}
