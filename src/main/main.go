package main

import (
	"config"
	"flag"
	"mitm"
)

func main() {

	conf := new(config.Cfg)
	conf.Port = flag.String("port", "8080", "Listen port")
	// conf.Addr = flag.String("addr", "127.0.0.1", "Listening  IP address")
	conf.Log = flag.String("log", "mitm.log", "Specify the log path")

	flag.Parse()

	ch := make(chan bool)
	mitm.Gomitmproxy(conf, ch)
	<-ch
}
