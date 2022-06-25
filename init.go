package gogadaddyclient

func NewGoDaddyClient(baseURL string, key string, secrect string) *GoaddyClient {

	var g GoaddyClient

	g.API = &GodaddyAPI{BaseURL: baseURL, Key: key, Secret: secrect}

	return &g
}
