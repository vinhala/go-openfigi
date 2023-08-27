package openfigi

type SearchJobResult struct {
	Data  []FigiResult `json:"data"`
	Error string       `json:"error,omitempty"`
	Next  string       `json:"next,omitempty"`
}
