package db

import "github.com/plancks-cloud/plancks-health/model"

//HealthCheck performs a health check and returns the state of the system
func HealthCheck() model.MessageOK {
	return model.Ok(true)
}
