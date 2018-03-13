package mitm

import (
	"config"
	"net/http"
	"time"
)

// Gomitmproxy create a mitm proxy and start it
func Gomitmproxy(conf *config.Cfg, ch chan bool) {
	handler := InitConfig(conf)
	server := &http.Server{
		Addr:         ":" + *conf.Port,
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
		Handler:      handler,
	}

	go func() {
		server.ListenAndServe()
		ch <- true
	}()

	return
}
