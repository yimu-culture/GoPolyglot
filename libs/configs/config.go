package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var GConfig *Config

func init() {
	GConfig = &Config{}
}

func InitConfigs(env string) (err error) {
	var (
		dir = "./configs/" + env + ".yml"
	)

	yamlConf, err := ioutil.ReadFile(dir)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(yamlConf, GConfig); err != nil {
		return err
	}

	return nil
}
