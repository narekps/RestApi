package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/narekps/RestApi/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file in .toml format")
}

func main() {
	// parse configs
	flag.Parse()

	log.Println("It works!")

	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Using default values. err:", err)
	}

	// server instance initialization
	apiServer := api.New(config)

	// server start
	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}
}
