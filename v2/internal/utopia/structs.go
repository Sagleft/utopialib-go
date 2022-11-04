package utopia

type UtopiaClient struct {
	data        Config
	logCallback LogCallback
}

type Config struct {
	// required
	Host   string `json:"host"`
	Token  string `json:"token"`
	Port   int    `json:"port"`
	WsPort int    `json:"wsport"`

	// optional
	Protocol              string      `json:"protocol"` // default: 127.0.0.1
	RequestTimeoutSeconds int         `json:"timeout"`
	Cb                    LogCallback `json:"-"`
}

// query is a filter for API requests
type query struct {
	Method  string                 `json:"method"`
	Token   string                 `json:"token"`
	Params  map[string]interface{} `json:"params"`
	Filters map[string]interface{} `json:"filter"`
}
