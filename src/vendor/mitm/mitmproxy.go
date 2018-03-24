package mitm

import (
	"config"
	"log"
	"net/http"
	"os"
	"time"
)

var logger *log.Logger

// Gomitmproxy create a mitm proxy and start it
func Gomitmproxy(conf *config.Cfg, ch chan bool) {
	tlsConfig := config.NewTLSConfig("ca-pk.pem", "ca-cert.pem", "", "")
	handler := InitConfig(conf, tlsConfig)
	server := &http.Server{
		Addr:         ":" + *conf.Port,
		ReadTimeout:  1 * time.Hour,
		WriteTimeout: 1 * time.Hour,
		Handler:      handler,
	}

	l, _ := os.Create(*conf.Log)
	logger = log.New(l, "[mitmproxy]", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logger.Println("Server is listening at ", server.Addr)

	go func() {
		server.ListenAndServe()
		ch <- true
	}()

	return
}
