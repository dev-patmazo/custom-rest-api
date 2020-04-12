package config

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/pelletier/go-toml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/yaml.v2"
)

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
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
	dbClient    *mongo.Database
)

func SetEnvConfig() {

	//Open config file
	cfg, cfgErr := filepath.Abs(configPath)
	if cfgErr != nil {
		log.Error(cfgErr)
	}

	//Read config file
	cfgCons, cfgConsErr := ioutil.ReadFile(cfg)
	if cfgConsErr != nil {
		log.Error(cfgConsErr)
	}

	//Parse config file content
	cfgRawErr := yaml.Unmarshal(cfgCons, &cfgMap)
	if cfgRawErr != nil {
		log.Error(cfgRawErr)
	}

	//Open environment file
	env, envErr := filepath.Abs(envPath)
	if envErr != nil {
		log.Error(envErr)
	}

	//Read environment file
	envCons, envConsErr := ioutil.ReadFile(env)
	if envConsErr != nil {
		log.Error(envConsErr)
	}

	//Parse environment file content
	envRawErr := toml.Unmarshal(envCons, &envMap)
	if envRawErr != nil {
		log.Error(envRawErr)
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

	log.Info("Setting up environment config...")
}

func SetDBCOnfig() {
	log.Info("Setting up database connection...")
	var dbPath = "mongodb://" + cfgTemplate.Database.Host + ":" + cfgTemplate.Database.Port

	client, clientErr := mongo.NewClient(options.Client().ApplyURI(dbPath))
	if clientErr != nil {
		log.Fatal(clientErr)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctxErr := client.Connect(ctx)
	if ctxErr != nil {
		log.Fatal(ctxErr)
	}
	defer client.Disconnect(ctx)

	pingErr := client.Ping(ctx, readpref.Primary())
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Info("Database connected!")
	dbClient = client.Database(cfgTemplate.Database.Dbname)

}

func SetLogConfig() {

	log.Info("Setting up logging service...")
	if envMap.Env == "prod" {
		log.SetLevel(log.InfoLevel)
		logrus.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetLevel(log.DebugLevel)
		logrus.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	}

}

func GetEnvInfo() (config ConfigTemplate) {
	return cfgTemplate
}

/*
	Use this function to start your database
	transaction. See : https://github.com/mongodb/mongo-go-driver
*/
func GetDBClient() *mongo.Database {
	return dbClient
}
