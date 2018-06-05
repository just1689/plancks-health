package model

type Contract struct {
	//Specification
	ImageAMD64       string `json:"imageAMD64" bson:"imageAMD64,omitempty"`
	RequiredMBMemory int    `json:"requiredMBMemory" bson:"requiredMBMemory,omitempty"`
	RequiredCPUCores int    `json:"requiredCPUCores" bson:"requiredCPUCores,omitempty"`
	ServiceName      string `json:"serviceName" bson:"serviceName,omitempty"`
}
