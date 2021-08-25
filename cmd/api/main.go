package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/bullterrier666/TestLogin/internal/app/api"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "configs/api.toml", "Path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("Starting...")
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find config file. Using default values: ", err)
	}
	//Прочитаем из .toml/.env
	server := api.New(config)
	log.Fatal(server.Start())
}
