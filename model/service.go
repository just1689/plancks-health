package model

type Service struct {
	ID             string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string `json:"name" bson:"name,omitempty"`
	Image          string `json:"image" bson:"image,omitempty"`
	HasWorked      bool   `json:"hasWorked" bson:"hasWorked,omitempty"`
	EffectiveDate  int64  `json:"effectiveDate" bson:"effectiveDate,omitempty"`
	RunUntil       int64  `json:"runUntil" bson:"runUntil,omitempty"`
	Network        string `json:"network" bson:"network,omitempty"`
	HealthyManaged bool   `json:"healthyManaged" bson:"healthyManaged,omitempty"`
	Replicas       int    `json:"replicas" bson:"replicas,omitempty"`
	ContractID     string `json:"contractId" bson:"contractId"`
}
