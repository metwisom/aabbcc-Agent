package pkg

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	Server  string    `yaml:"server"`
	Aspects []*Aspect `yaml:"aspects"`
}

func (c *Config) GetConf() *Config {

	yamlFile, err := os.ReadFile("aspect.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
