package mindcli

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Robotport struct {
	RobotIp         string
	ServeMPKPort    int
	ServeRemotePort int
}

type PortConfig struct {
	Robotsport []Robotport `json:"Robotsport"`
	path       string
}

func NewPortConfig(path string) *PortConfig {
	var config PortConfig
	file, err := os.Open(path)
	if err != nil {
		return &PortConfig{path: path}
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return &PortConfig{path: path}
	}
	config.path = path
	return &config
}

func (config *PortConfig) Write() error {
	configJson, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(config.path, configJson, 0644)
}
