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
	logger.Info("test log written by zap.Logger")
}

func TestSetLogLevel(t *testing.T) {
	logger, err := SetLogLevel("debug")
	logger.Debug("here is the sample log written with debug level")
	assert.NotNil(t, logger)
	assert.Nil(t, err)
}

func TestSetLogLevelInvalidLevel(t *testing.T) {
	logger, err := SetLogLevel("debugggg")
	assert.Nil(t, logger)
	assert.NotNil(t, err)
}

func TestSetLogLevelNotInitialized(t *testing.T) {
	logger = nil
	logger, err := SetLogLevel("adsfadsfadsf")
	assert.Nil(t, logger)
	assert.NotNil(t, err)
}
