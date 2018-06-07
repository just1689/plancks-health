package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRequest(t *testing.T) {

	url := "https://ifconfig.co"
	bytes, err := GetRequest(url)

	assert.NotNil(t, bytes, "The website should return some bytes")
	assert.Nil(t, err, "The err should be nil")

}
