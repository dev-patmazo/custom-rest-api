package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Username string `yaml:"user"`
	Password string `yaml:"pass"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

type Config struct {
	Dev struct {
		Active   string
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"dev"`

	Qa struct {
		Active   string
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"qa"`

	Stg struct {
		Active   string
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"stg"`

	Prod struct {
		Active   string
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"prod"`
}

type ConfigTemplate struct {
	Active   string
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
}

var cfg Config
var cfgTemplate ConfigTemplate
var basePath = "../rest-api/config/config.yaml"

func InitializeConfig() (config ConfigTemplate) {

	//Open file
	file, err := filepath.Abs(basePath)
	if err != nil {
		processError(err)
	}

	//Read file
	fileContent, errFile := ioutil.ReadFile(file)
	if errFile != nil {
		processError(errFile)
	}

	//Parse file content
	parseErr := yaml.Unmarshal(fileContent, &cfg)
	if parseErr != nil {
		processError(err)
	}

	//Check environment
	if cfg.Dev.Active == "true" {
		cfgTemplate = cfg.Dev
		return cfgTemplate
	}
	if cfg.Qa.Active == "true" {
		cfgTemplate = cfg.Qa
		return cfgTemplate
	}
	if cfg.Stg.Active == "true" {
		cfgTemplate = cfg.Stg
		return cfgTemplate
	}
	if cfg.Prod.Active == "true" {
		cfgTemplate = cfg.Prod
		return cfgTemplate
	}

	return cfgTemplate
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
