package main

import (
	"config"
	"flag"
)

func main() {
	conf := new(config.Cfg)
	conf.Port = flag.String("port", "8080", "Listen port")
	conf.Addr = flag.String("addr", "127.0.0.1", "Listening  IP address")

	flag.Parse()

}
