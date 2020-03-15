package main

import (
	neo "github.com/minskylab/neocortex"

	"github.com/pkg/errors"
)

func main() {
	config, err := extractConfigFromEnv()
	if err != nil {
		panic(errors.Cause(err))
	}

	brain, err := loadCognitive(config)
	if err != nil {
		panic(errors.Cause(err))
	}

	channels, err := loadChannels(config)
	if err != nil {
		panic(errors.Cause(err))
	}
	fb := channels[0]

	engine, err := neo.Default(nil, brain, channels...)
	if err != nil {
		panic(errors.Cause(err))
	}

	engine.ResolveAny(fb, func(c *neo.Context, in *neo.Input, out *neo.Output, response neo.OutputResponse) error {
		return response(c, out)
	})

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
