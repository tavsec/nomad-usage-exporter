package main

import "github.com/hashicorp/nomad/api"

type NomadApi struct {
	config      Config
	nomadClient *api.Client
}

func NewNomadApi(config Config) (NomadApi, error) {
	nomadApi := NomadApi{config: config}
	nomadConfig := &api.Config{
		Address:   config.Host,
		Region:    config.Region,
		Namespace: config.Namespace,
		SecretID:  config.Token,
	}

	var err error
	nomadApi.nomadClient, err = api.NewClient(nomadConfig)
	if err != nil {
		return nomadApi, err
	}
	return nomadApi, nil
}

func (nomadApi NomadApi) fetchJobs() *api.Jobs {
	return nomadApi.nomadClient.Jobs()
}

func (nomadApi NomadApi) fetchDeployments(jobId string, all bool) ([]*api.Deployment, error) {
	deployments, _, err := nomadApi.nomadClient.Jobs().Deployments(jobId, all, nil)
	return deployments, err
}

func (nomadApi NomadApi) fetchJobVersions(jobId string, diffs bool) ([]*api.Job, error) {
	versions, _, _, err := nomadApi.nomadClient.Jobs().Versions(jobId, diffs, nil)
	return versions, err
}

func (nomadApi NomadApi) fetchJobAllocations(jobId string, all bool) ([]*api.AllocationListStub, error) {
	allocations, _, err := nomadApi.nomadClient.Jobs().Allocations(jobId, all, nil)
	return allocations, err
}
