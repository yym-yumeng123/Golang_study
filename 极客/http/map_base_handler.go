package http

import "net/http"

type Routable interface {
	Route(method string, pattern string, handleFunc HandlerFunc)
}

type Handler interface {
	ServeHTTP(c *Context)
	Routable
}

// HandlerBasedOnMap 实现了 Handle 接口
type HandlerBasedOnMap struct {
	// key 应该是 method + url
	handlers map[string]HandlerFunc
}

func (h *HandlerBasedOnMap) Route(
	method string,
	pattern string,
	handleFunc HandlerFunc) {
	key := h.Key(method, pattern)
	h.handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	key := h.Key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		handler(c)
	} else {
		c.W.WriteHeader(http.StatusNotFound)
	}
}

func (h *HandlerBasedOnMap) Key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{handlers: make(map[string]HandlerFunc)}
}
