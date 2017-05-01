package response

// Result ma man
type Result struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}
