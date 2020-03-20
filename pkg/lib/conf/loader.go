package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error to read file for path: %s", path)
	}

	if err := yaml.Unmarshal(yamlFile, RuntimeConf); err != nil {
		return fmt.Errorf("error to unmarshal yaml to Config")
	}

	return nil
}
