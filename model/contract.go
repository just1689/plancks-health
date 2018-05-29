package model

type Contract struct {
	//Audit & admin
	// ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	ID        string `json:"_id,omitempty" bson:"_id,omitempty"`
	Account   string `json:"account" bson:"account,omitempty"`
	Signature string `json:"signature" bson:"signature,omitempty"`
	Timestamp int64  `json:"timestamp" bson:"timestamp,omitempty"`

	//Specification
	ImageAMD64       string `json:"imageAMD64" bson:"imageAMD64,omitempty"`
	Instances        int    `json:"instances" bson:"instances,omitempty"`
	Replicas         int    `json:"replicas" bson:"replicas,omitempty"`
	RequiredMBMemory int    `json:"requiredMBMemory" bson:"requiredMBMemory,omitempty"`
	RequiredCPUCores int    `json:"requiredCPUCores" bson:"requiredCPUCores,omitempty"`
	SecondsToLive    int64  `json:"secondsToLive" bson:"secondsToLive,omitempty"`
	AllowSuicide     bool   `json:"allowSuicide" bson:"allowSuicide,omitempty"`
	StartStrategy    string `json:"startStrategy" bson:"startStrategy,omitempty"`
	ServiceName      string `json:"serviceName" bson:"serviceName,omitempty"`
}
