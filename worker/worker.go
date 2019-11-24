package worker

import "github.com/kanapuli/eks-ami-update/config"

//Run reads the config params and updates ami changes
func Run(cfg config.Configuration) error {
	asgDesiredCapacity := getAsgDesiredCapacity(cfg.asgName)
	if checkAsgHealth(cfg.asgName, cfg.asgDesiredCapacity) {
	}
}

func checkAsgHealth() bool {
}

func getAsgDesiredCapacity() (int, error) {
}
