package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
)

func extractConfigFromEnv() (*config, error) {
	config := new(config)
	if err := env.Parse(config); err != nil {
		return nil, errors.Wrap(err, "environment variables failed at try to parsing")
	}
	return config, nil
}
