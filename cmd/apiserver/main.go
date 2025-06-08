package main

import (
	"ShorterAPI/internal/app/apiserver"
	"ShorterAPI/internal/repository/postgres"
	"flag"
	"github.com/BurntSushi/toml"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "path", "config/apiserver.toml", "path to config")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()

	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}

	pool, err := postgres.InitConnectionPool(config.Db.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config, pool)

	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
