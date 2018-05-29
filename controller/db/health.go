package db

import "git.amabanana.com/plancks-cloud/pc-go-brutus/model"

//HealthCheck performs a health check and returns the state of the system
func HealthCheck() model.MessageOK {
	return model.Ok(true)
}
