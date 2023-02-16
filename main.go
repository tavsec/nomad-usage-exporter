package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"sync"
)

var config Config
var nomadApi NomadApi
var wg sync.WaitGroup

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
		wg.Add(1)
		go displayVersions(job.ID)
	}
	wg.Wait()

}

func displayVersions(jobId string) {
	var log = log.New().WithField("jobId", jobId)
	versions, err := nomadApi.fetchJobVersions(jobId, false)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	for _, version := range versions {
		log.Println(fmt.Sprintf("Fetched version %d, which was deployed on %d", *version.Version, *version.SubmitTime))
		log = log.WithField("version", *version.Version)
		log.Println("Fetching resource usage for all tasks in all task groups")
	}

	defer wg.Done()
}
