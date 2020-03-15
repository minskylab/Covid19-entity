package main

type config struct {
	FbAccessToken  string `env:"FB_ACCESS_TOKEN"`
	FbVerifySecret string `env:"FB_VERIFY_SECRET"`
	FbPageID       string `env:"FB_PAGE_ID"`

	WnURL         string `env:"WN_URL"`
	WnVersion     string `env:"WN_VERSION"`
	WnAssistantID string `env:"WN_ASSISTANT_ID"`
	WnAPIKey      string `env:"WN_API_KEY"`
}
