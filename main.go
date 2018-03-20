package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
)

var (
	Service string
	Version string
	Commit  string
)

func main() {
	app := cli.NewApp()
	app.Name = Service
	app.Usage = "An image resizer microservice"
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "configuration file",
			Value:  "config.yaml",
			EnvVar: "CDN_ORIGIN_CONFIG",
		},
	}
	app.Before = configure
	app.Action = action

	app.Run(os.Args)
}

func action(c *cli.Context) error {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", RootHandler)
	r.HandleFunc("/health", HealthHandler)
	r.HandleFunc("/{profile}/{source}", RequestHandler)
	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, r))
	return nil
}
