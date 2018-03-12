package main

import (
	"config"
	"flag"
	"fmt"
	"mitm"
	"mylog"
	"os"
)

func main() {

	conf := new(config.Cfg)
	conf.Port = flag.String("port", "8080", "Listen port")
	// conf.Addr = flag.String("addr", "127.0.0.1", "Listening  IP address")
	conf.Log = flag.String("log", "mitm.log", "Specify the log path")

	flag.Parse()

	log, err := os.Create(*conf.Log)
	if err != nil {
		fmt.Println("fail to create log file " + err.Error())
	}
	mylog.SetLog(log)
	mitm.Gomitmproxy(conf)
}
