package utopia

type UtopiaClient struct {
	// required, public
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Token    string `json:"token"`
	Port     int    `json:"port"`

	// optional, public
	RequestTimeoutSeconds int `json:"timeout"`
	WsPort                int `json:"wsport"`

	// protected
	logCallback LogCallback
}

// query is a filter for API requests
type query struct {
	Method  string                 `json:"method"`
	Token   string                 `json:"token"`
	Params  map[string]interface{} `json:"params"`
	Filters map[string]interface{} `json:"filter"`
}
