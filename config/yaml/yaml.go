package yaml

import (
	yamlService "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Server struct {
		Port string `yaml:"port"`
	}
	DataSource struct {
		Connection string `yaml:"connection"`
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		Password   string `yaml:"password"`
		Username   string `yaml:"username"`
		Database   string `yaml:"database"`
	}

	Redis struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}
}

func Conf() *Yaml {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("app.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yamlService.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return conf
}
