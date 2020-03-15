package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/channels/facebook"
	"github.com/pkg/errors"
)

func loadChannels(config *config) ([]neo.CommunicationChannel, error) {
	channels := make([]neo.CommunicationChannel, 0)

	fb, err := facebook.NewChannel(facebook.ChannelOptions{
		AccessToken: config.FbAccessToken,
		VerifyToken: config.FbVerifySecret,
		PageID:      config.FbPageID,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error at create new facebook channel")
	}

	channels = append(channels, fb)

	return channels, nil
}
