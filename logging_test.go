package commons

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	t.Log("getting logger")
	logger := GetLogger()
	assert.NotNil(t, logger)
	t.Log("will try logger for debugging")
	logger.Log("info", "test log written by VpnbeastLogger")
}

func TestGetLoggerInvalidLogLevel(t *testing.T) {
	t.Log("getting logger")
	logger := GetLogger()
	assert.NotNil(t, logger)
	t.Log("will try logger for debugging")
	logger.Log("infoooo", "test log written by VpnbeastLogger")
}
