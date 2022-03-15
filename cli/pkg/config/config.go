package config

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

type Config struct {
	Name     string           `yaml:"name"`
	CSI      string           `yaml:"csi"`
	Pipeline v1beta1.Pipeline `yaml:"pipeline"`
}
