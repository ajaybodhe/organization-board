package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	configFile = "./resource/organization-board.cfg"
)

var (
	_config *ApiConfig
)

// ApiConfig : Sqllite connection config
type ApiConfig struct {
	Host   string     `json:"host"`
	Port   int        `json:"port"`
	SQLite *sqliteCfg `json:"sqlite"`
}

type sqliteCfg struct {
	DataSourceName string `json:"data_source_file"`
}

func (api *ApiConfig) String() string {
	byts, _ := json.Marshal(api)
	return string(byts)
}

func init() {
	file, err := ioutil.ReadFile(configFile)
	if nil != err {
		log.Printf("Error while reading config file: %s:%s", configFile, err.Error())
		os.Exit(1)
	}

	_config = new(ApiConfig)
	err = json.Unmarshal(file, _config)
	if nil != err {
		log.Printf("Error while reading configuration: %s:%s", configFile, err.Error())
		os.Exit(1)
	}

	// if byts, err := json.Marshal(_config); nil != err {
	// 	log.Fatalf("Error while json.Marshal config structure: %s", err.Error())
	// } else {
	// 	log.Println("Application Configuration is:", string(byts))
	// }
}

func Config() *ApiConfig {
	return _config
}
