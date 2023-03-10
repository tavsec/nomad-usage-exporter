package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"sync"
)

var config Config
var nomadApi NomadApi
var wg sync.WaitGroup
var resourceUsages []ResourceUsage

func init() {
	err := envconfig.Process("nomad-exporter", &config)
	if err != nil {
		log.Fatalln("Error reading config: " + err.Error())
	}

	nomadApi, err = NewNomadApi(config)
	if err != nil {
		log.Fatalln("Error creating Nomad API client: " + err.Error())
	}

	InitDynamoDb()

	err = godotenv.Load()
	if err != nil {
		log.Info("Error loading .env file - fallback to ENVIRONMENT VARIABLES")
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

	allocations, err := nomadApi.fetchJobAllocations(jobId, false)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	taskAllocations := make([]Allocation, 0)
	for _, allocation := range allocations {
		taskAllocations = append(taskAllocations, Allocation{
			ID:       allocation.ID,
			NodeID:   allocation.NodeID,
			NodeName: allocation.NodeName,
		})
	}

	for _, version := range versions {
		log.Println(fmt.Sprintf("Fetched version %d, which was deployed on %d", *version.Version, *version.SubmitTime))
		log = log.WithField("version", *version.Version)
		log.Println("Fetching resource usage for all tasks in all task groups")
		for _, taskGroup := range version.TaskGroups {
			log = log.WithField("taskGroup", *taskGroup.Name)
			for _, task := range taskGroup.Tasks {
				log = log.WithField("task", task.Name)
				resourceUsage := ResourceUsage{
					ID:                fmt.Sprintf("%s-%d", jobId, *version.Version),
					TaskName:          task.Name,
					JobId:             jobId,
					CPUPerInstance:    *task.Resources.CPU,
					MemoryPerInstance: *task.Resources.MemoryMB,
					NumberOfInstances: *taskGroup.Count,
					CPUTotal:          *taskGroup.Count * *task.Resources.CPU,
					MemoryTotal:       *taskGroup.Count * *task.Resources.MemoryMB,
					ChangedAt:         *version.SubmitTime,
					VersionId:         *version.Version,
					Allocations:       taskAllocations,
				}

				err := StoreResourceUsage(resourceUsage)
				if err != nil {
					log.Fatalln("Error storing resource usage to DynamoDB: " + err.Error())
				}

				resourceUsages = append(resourceUsages, resourceUsage)
			}
		}
	}

	defer wg.Done()
}
