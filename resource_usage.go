package main

type ResourceUsage struct {
	TaskName          string
	JobId             string
	CPUPerInstance    int
	MemoryPerInstance int
	NumberOfInstances int
	CPUTotal          int
	MemoryTotal       int
	ChangedAt         int64
}
