package utopia

import (
	"github.com/Sagleft/utopialib-go/v2/internal/reqhandler"
	"github.com/beefsack/go-rate"
)

type UtopiaClient struct {
	reqHandler  reqhandler.RequestHandler
	data        Config
	logCallback LogCallback
	limiters    rateLimiters
}

type rateLimiters map[string]*rate.RateLimiter

type Config struct {
	// required
	Host   string `json:"host" yaml:"host"` // default: 127.0.0.1
	Token  string `json:"token" yaml:"token"`
	Port   int    `json:"port" yaml:"port"`
	WsPort int    `json:"wsport" yaml:"wsport"`

	// optional
	Protocol              string      `json:"protocol" yaml:"protocol"` // default: http
	RequestTimeoutSeconds int         `json:"timeout" yaml:"timeout"`
	Cb                    LogCallback `json:"-" yaml:"-"`
}

// query is a filter for API requests
type query struct {
	Method  string                 `json:"method"`
	Token   string                 `json:"token"`
	Params  map[string]interface{} `json:"params"`
	Filters map[string]interface{} `json:"filter"`
}
