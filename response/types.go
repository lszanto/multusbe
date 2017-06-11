package response

// Result ma man
type Result struct {
	Result interface{} `json:"result,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

// LoginResult holds a JWT token
type LoginResult struct {
	Token string `json:"token,omitempty"`
}
