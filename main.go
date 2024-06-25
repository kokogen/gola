package main

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/gola/internal/apiserver"
	"github.com/gola/internal/conf"
	"github.com/sirupsen/logrus"
)

const configPath = "apiserver.yaml"

func readConfig(_configPath string) (*conf.Config, error) {
	var config conf.Config = conf.Config{}

	file, err := os.Open(_configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	logrus.Info("apiserver started")

	config, err := readConfig(configPath)
	if err != nil {
		panic(err)
	}

	if err := apiserver.Start(config); err != nil {
		panic(err)
	}

}
