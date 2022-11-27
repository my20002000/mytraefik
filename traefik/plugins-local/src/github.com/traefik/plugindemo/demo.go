package plugindemo

import (
	"context"
	"net/http"
	"strings"
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
	strSplit := strings.Split(req.RemoteAddr, ":")
	idx := len(strSplit)
	if idx > 1 {
		req.Header.Set("X-Real-Port", strSplit[idx-1])
	}
	a.next.ServeHTTP(rw, req)
}

