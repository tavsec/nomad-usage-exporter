package main

type Config struct {
	Host      string `envconfig:"NOMAD_ADDR" default:"http://localhost:4646"`
	Region    string `envconfig:"NOMAD_REGION" default:""`
	Namespace string `envconfig:"NOMAD_NAMESPACE" default:""`
	Token     string `envconfig:"NOMAD_TOKEN" default:""`
}
