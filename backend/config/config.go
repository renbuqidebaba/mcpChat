package config

type LLMConfig struct {
	BaseUrl string
	ApiKey  string
}

func NewLLMConfig(baseUrl, apiKey string) *LLMConfig {
	return &LLMConfig{
		BaseUrl: baseUrl,
		ApiKey:  apiKey,
	}
}
