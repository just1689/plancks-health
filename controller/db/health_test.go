package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthCheck(t *testing.T) {

	r := HealthCheck()
	assert.NotNil(t, r, "The health check should return a non-nil struct")
	assert.True(t, r.Ok, "The health check OK bool should be true")

}
