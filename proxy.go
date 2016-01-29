package ginproxy

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

type (
	ProxyRouter struct {
		upstream string
		proxy    *httputil.ReverseProxy
	}
)

func NewGinProxy(upstream string) (*ProxyRouter, error) {
	u, err := url.Parse(upstream)
	if err != nil {
		return nil, err
	}

	return &ProxyRouter{
		proxy: httputil.NewSingleHostReverseProxy(u),
	}, nil
}

func (p *ProxyRouter) Handler(c *gin.Context) {
	p.proxy.ServeHTTP(c.Writer, c.Request)
}
