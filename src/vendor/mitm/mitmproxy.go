package mitm

import (
	"config"
	"net/http"
	"time"
)

// Gomitmproxy create a mitm proxy and start it
func Gomitmproxy(conf *config.Cfg) {
	server := &http.Server{
		Addr:         ":" + *conf.Port,
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
	}

	go func() {
		server.ListenAndServe()
	}()

	return
}
