package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

var config Config

func init() {
	err := envconfig.Process("nomad-exporter", &config)
	if err != nil {
		log.Fatalln("Error reading config: " + err.Error())
	}
}

func main() {

}
