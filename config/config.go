package config

//Configuration is  the application configuration
type Configuration struct {
	AutoscalingEnabled  bool   `json:"autoscalingEnabled"`
	AutoscalerNamespace string `json:"autoscalerNamespace"`
	AsgName             string `json:"asgName"`
	AsgDesiredCapacity  int    `json:"asgDesiredCapacity"`
	DryRun              bool   `json:"dryRun"`
}
