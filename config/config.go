package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
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
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"dev"`

	Qa struct {
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"qa"`

	Stg struct {
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"stg"`

	Prod struct {
		Server   *Server   `yaml:"server"`
		Database *Database `yaml:"database"`
	} `yaml:"prod"`
}

type ConfigTemplate struct {
	Server   *Server   `yaml:"server"`
	Database *Database `yaml:"database"`
}

type Environment struct {
	Env string
}

var (
	cfgMap      Config
	cfgTemplate ConfigTemplate
	envMap      Environment
	basePath    = "../rest-api/"
	configPath  = basePath + "config/config.yaml"
	envPath     = basePath + ".air.conf"
)

func InitializeConfig() (config ConfigTemplate) {

	//Open config file
	cfg, cfgErr := filepath.Abs(configPath)
	if cfgErr != nil {
		processError(cfgErr)
	}

	//Read config file
	cfgCons, cfgConsErr := ioutil.ReadFile(cfg)
	if cfgConsErr != nil {
		processError(cfgConsErr)
	}

	//Parse config file content
	cfgRawErr := yaml.Unmarshal(cfgCons, &cfgMap)
	if cfgRawErr != nil {
		processError(cfgRawErr)
	}

	//Open environment file
	env, envErr := filepath.Abs(envPath)
	if envErr != nil {
		processError(envErr)
	}

	//Read environment file
	envCons, envConsErr := ioutil.ReadFile(env)
	if envConsErr != nil {
		processError(envConsErr)
	}

	//Parse environment file content
	envRawErr := toml.Unmarshal(envCons, &envMap)
	if envRawErr != nil {
		processError(envRawErr)
	}

	if envMap.Env == "dev" {
		cfgTemplate = cfgMap.Dev
		return cfgTemplate
	}
	if envMap.Env == "qa" {
		cfgTemplate = cfgMap.Qa
		return cfgTemplate
	}
	if envMap.Env == "stg" {
		cfgTemplate = cfgMap.Stg
		return cfgTemplate
	}
	if envMap.Env == "prod" {
		cfgTemplate = cfgMap.Prod
		return cfgTemplate
	}

	return cfgTemplate
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}
