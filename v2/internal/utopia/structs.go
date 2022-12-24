package utopia

import (
	"net/http"

	"github.com/beefsack/go-rate"
)

type UtopiaClient struct {
	httpClient  *http.Client
	data        Config
	logCallback LogCallback
	limiters    rateLimiters
}

type rateLimiters map[string]*rate.RateLimiter

type Config struct {
	// required
	Host   string `json:"host"` // default: 127.0.0.1
	Token  string `json:"token"`
	Port   int    `json:"port"`
	WsPort int    `json:"wsport"`

	// optional
	Protocol              string      `json:"protocol"` // default: http
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
