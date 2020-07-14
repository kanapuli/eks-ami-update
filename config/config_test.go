package config

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	conf := New()
	if !reflect.DeepEqual(new(Configuration), conf) {
		t.Errorf("Expected %v but got %v", new(Configuration), conf)
	}
}

func TestRead(t *testing.T) {
	jsonInput := strings.NewReader("{ \"autoscalingEnabled\": true, \"autoscalerNamespace\": \"default\", \"asgName\": \"eks-1-16-asg\" , \"asgDesiredCapacity\": 2, \"dryRun\": false }")
	conf := new(Configuration)
	err := conf.Read(jsonInput)
	if err != nil {
		t.Errorf("Expected Json parsing but got %v", err)
	}
	if conf.AsgDesiredCapacity != 2 || conf.AutoscalingEnabled != true || conf.DryRun != false || conf.AutoscalerNamespace != "default" || conf.AsgName != "eks-1-16-asg" {
		t.Errorf("Json parsing values expected: %v but got %v", jsonInput, conf)
	}
}
