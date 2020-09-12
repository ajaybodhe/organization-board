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
}

func Config() *ApiConfig {
	return _config
}
