package ssfrepo

import (
	"os"
	"path/filepath"

	"github.com/thesecuresoftwarefactory/ssf/cli/pkg/config"
	"gopkg.in/yaml.v2"
)

func Create(configFile string) error {
	filename, _ := filepath.Abs(configFile)
	yamlFile, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	var config config.Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil

}
