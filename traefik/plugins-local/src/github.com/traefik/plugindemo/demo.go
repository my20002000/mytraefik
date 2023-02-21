package plugindemo

import (
	"context"
	"net/http"
	"strings"
	"net"
)

// Config the plugin configuration.
type Config struct {
	Port bool
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Port: true,
	}
}

// Demo a Demo plugin.
type Demo struct {
	next   http.Handler
	config *Config
	name   string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Demo{
		config: config,
		next:   next,
		name:   name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	host, port, err := net.SplitHostPort(req.RemoteAddr)
	if err != nil {
		req.Header.Set("X-Real-Port", "0")
	}else{
		req.Header.Set("X-Real-Port", port)
	}
	a.next.ServeHTTP(rw, req)
}

