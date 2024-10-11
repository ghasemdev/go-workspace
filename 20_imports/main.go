package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"os"
)

type Config struct {
	Login struct {
		User     string
		Password string
	}
}

func main() {
	file, err := os.Open("config.toml")
	if err != nil {
		log.Fatalf("Error cam't open config file - %s\n", err)
	}

	defer file.Close()

	ctg := &Config{}
	dec := toml.NewDecoder(file)
	if err := dec.Decode(ctg); err != nil {
		log.Fatalf("Error: can't decode config file - %s\n", err)
	}

	fmt.Printf("%+v", ctg)
}
