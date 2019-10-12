package hatchbuckapi

// HatchbuckClient is the api client for the Hatchbuck API
type HatchbuckClient struct {
	key     string
	baseURL string
}

// Init creates an API client and returns it
func Init(key string) HatchbuckClient {
	hb := HatchbuckClient{
		key:     key,
		baseURL: "https://api.hatchbuck.com/api/v1",
	}
	return hb
}
