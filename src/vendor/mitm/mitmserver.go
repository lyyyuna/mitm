package mitm

import (
	"config"
	"net/http"
)

// HandlerWrapper wrapper of handler for http server
type HandlerWrapper struct {
	Config  *config.Cfg
	wrapped http.Handler
}

// InitConfig init HandlerWrapper
func InitConfig(conf *config.Cfg) *HandlerWrapper {
	handler := &HandlerWrapper{
		Config: conf,
	}

	return handler
}

// ServeHTTP the main function interface for http handler
func (handler *HandlerWrapper) ServeHTTP(resp http.ResponseWriter, req http.Request) {
	if req.Method == "CONNECT" {

	} else {

	}
}
