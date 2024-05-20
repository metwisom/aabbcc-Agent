package pkg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Server  string     `yaml:"server"`
	Token   string     `yaml:"token"`
	Aspects AspectList `yaml:"aspect"`
}

func (c *Config) GetConf() *Config {

	if yamlFile, err := os.ReadFile("aspect.yaml"); err != nil {
		panic(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	} else {
		err := yaml.Unmarshal(yamlFile, c)
		if err != nil {
			panic(fmt.Sprintf("Unmarshal: %v", err))
		}
	}

	return c
}
