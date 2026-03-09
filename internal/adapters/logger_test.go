// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package adapters

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	confmocks "github.hpe.com/cloud/go-gadgets/x/config/mocks"
)

func TestProvideLogger_HappyPath(t *testing.T) {
	conf := &confmocks.Config{}

	conf.On("Get", "logging.loglevel").Return("debug", nil)

	logger, err := ProvideLogger(conf)
	assert.NoError(t, err)
	assert.NotNil(t, logger)
}

func TestProvideLogger_ConfigReturnsError(t *testing.T) {
	conf := &confmocks.Config{}

	conf.On("Get", "logging.loglevel").Return(nil, errors.New("test errror"))

	logger, err := ProvideLogger(conf)
	assert.EqualError(t, err, "config was not found with key: logging.loglevel")
	assert.Nil(t, logger)
}
