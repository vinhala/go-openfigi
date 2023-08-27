package openfigi

type FilterJobResult struct {
	Data  []FigiResult `json:"data"`
	Error string       `json:"error,omitempty"`
	Next  string       `json:"next,omitempty"`
	Total int          `json:"total,omitempty"`
}
