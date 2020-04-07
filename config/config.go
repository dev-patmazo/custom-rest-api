package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

var logger = Logger()

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

func init() {

	logrus.SetLevel(logrus.WarnLevel)

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
	}
	if envMap.Env == "qa" {
		cfgTemplate = cfgMap.Qa
	}
	if envMap.Env == "stg" {
		cfgTemplate = cfgMap.Stg
	}
	if envMap.Env == "prod" {
		cfgTemplate = cfgMap.Prod
	}

}

func InitializeConfig() (config ConfigTemplate) {

	return cfgTemplate

}

func processError(err error) {
	logger.Error(err)
	os.Exit(2)
}

func Logger() *logrus.Logger {

	if envMap.Env == "prod" {
		log := logrus.New()
		log.SetLevel(logrus.PanicLevel)
		return log
	}

	return logrus.New()

}
