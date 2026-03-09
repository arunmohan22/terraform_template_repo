// (C) Copyright 2023 Hewlett Packard Enterprise Development LP

package adapters

import (
	"github.hpe.com/cloud/go-gadgets/x/config"
	"github.hpe.com/cloud/go-gadgets/x/logging"
)


// ProvideLogger provides a go-gadgets JSON logger
func ProvideLogger(conf config.Config) (logging.Logger, error) {
	loglevel, err := config.Get[string](conf, logLevelKey)
	if err != nil {
		return nil, err
	}

	return logging.NewZapJSONLogger(loglevel)
}
