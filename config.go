package main

import (
	"io/ioutil"
	"log"
	"runtime"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

var config *Settings

type Settings struct {
	Port            int                          `port`
	ImaginHost      string                       `imaginhost`
	ImaginPort      int                          `imaginport`
	EnableUrlSource bool                         `enableurlsource`
	Profiles        map[string]map[string]string `profiles`
}

func configure(c *cli.Context) error {
	runtime.GOMAXPROCS(runtime.NumCPU())
	file := c.String("config")
	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("could not read file")
		return err
	}

	// default settings
	cfg := Settings{
		Port:       8080,
		ImaginHost: "127.0.0.1",
		ImaginPort: 9000,
	}

	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		log.Fatal("could not parse yaml file")
		return err
	}

	config = &cfg

	return nil
}
