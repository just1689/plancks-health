package model

type Contract struct {
	//Specification
	ImageAMD64       string `json:"imageAMD64"`
	RequiredMBMemory int    `json:"requiredMBMemory"`
	RequiredCPUCores int    `json:"requiredCPUCores"`
	ServiceName      string `json:"serviceName"`
}
