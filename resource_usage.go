package main

type ResourceUsage struct {
	ID                string
	TaskName          string
	JobId             string
	VersionId         uint64
	CPUPerInstance    int
	MemoryPerInstance int
	NumberOfInstances int
	CPUTotal          int
	MemoryTotal       int
	ChangedAt         int64
}
