package main

import (
	neo "github.com/minskylab/neocortex"
	"github.com/minskylab/neocortex/cognitive/watson"
)

func loadCognitive(config *config) (neo.CognitiveService, error) {
	return watson.NewCognitive(watson.NewCognitiveParams{
		Url:         config.WnURL,
		Version:     config.WnVersion,
		AssistantID: config.WnAssistantID,
		ApiKey:      config.WnAPIKey,
	})
}
