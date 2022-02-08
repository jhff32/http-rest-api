package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jhff32/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configs-path", "configs/apiserver.toml", "path to configs file")
}

func main() {
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
