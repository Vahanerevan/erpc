package erpc

type Config struct {
	Secret string `json:"secret"`
	URL    string `json:"url"`
	Debug  *bool  `json:"debug"`
}
