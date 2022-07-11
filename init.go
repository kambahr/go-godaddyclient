package godaddyclient

func NewGoDaddyClient(baseURL string, version string, key string, secrect string) *GoaddyClient {

	var g GoaddyClient

	g.api = &GodaddyAPI{BaseURL: baseURL, Version: version, Key: key, Secret: secrect}

	g.Domain = &Domain{&GodaddyAPI{BaseURL: baseURL, Version: version, Key: key, Secret: secrect}}

	return &g
}
