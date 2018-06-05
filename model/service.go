package model

type Service struct {
	Name             string `json:"name"`
	Image            string `json:"image"`
	Network          string `json:"network"`
	HealthyManaged   bool   `json:"healthyManaged"`
	RequiredMBMemory int    `json:"requiredMBMemory"`
}
