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
	Host   string `json:"host" yaml:"host" envconfig:"UTOPIA_HOST" default:"127.0.0.1"`
	Token  string `json:"token" yaml:"token" envconfig:"UTOPIA_TOKEN"`
	Port   int    `json:"port" yaml:"port" envconfig:"UTOPIA_PORT" default:"22825"`
	WsPort int    `json:"wsport" yaml:"wsport" envconfig:"UTOPIA_WS_PORT" default:"25000"`

	// optional
	Protocol              string      `json:"protocol" yaml:"protocol" envconfig:"UTOPIA_PROTO" default:"http"`
	RequestTimeoutSeconds int         `json:"timeout" yaml:"timeout" envconfig:"UTOPIA_CONN_TIMEOUT" default:"5000"`
	Cb                    LogCallback `json:"-" yaml:"-"`
}

// query is a filter for API requests
type query struct {
	Method  string                 `json:"method"`
	Token   string                 `json:"token"`
	Params  map[string]interface{} `json:"params"`
	Filters map[string]interface{} `json:"filter"`
}
