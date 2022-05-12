package utopiago

// Query is a filter for API requests
type Query struct {
	Method string                 `json:"method"`
	Token  string                 `json:"token"`
	Params map[string]interface{} `json:"params"`
}

// UtopiaClient lets you connect to Utopia Client
type UtopiaClient struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Token    string `json:"token"`
	Port     int    `json:"port"`
}
