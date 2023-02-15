package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

var config Config
var nomadApi NomadApi

func init() {
	err := envconfig.Process("nomad-exporter", &config)
	if err != nil {
		log.Fatalln("Error reading config: " + err.Error())
	}

	nomadApi, err = NewNomadApi(config)
	if err != nil {
		log.Fatalln("Error creating Nomad API client: " + err.Error())
	}
}

func main() {
	jobs, _, err := nomadApi.fetchJobs().List(nil)
	if err != nil {
		log.Fatalln("Error fetching all jobs: " + err.Error())
	}
	for _, job := range jobs {
		println(job.Name)
	}
}
