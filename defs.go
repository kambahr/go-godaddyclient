package gogadaddyclient

type Result struct {
	StatusCode int
	Message    string
}

type GoaddyClient struct {
	API IGodaddyAPI
}

type GodaddyAPI struct {
	BaseURL string
	Key     string
	Secret  string
}
