package erpc

type RequestConfig struct {
	Secret string `json:"secret"`
	URL    string `json:"url"`
}
