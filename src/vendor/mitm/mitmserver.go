package mitm

import (
	"bufio"
	"config"
	"io"
	"mylog"
	"net"
	"net/http"
	"net/http/httputil"
	"regexp"
	"time"
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
func (handler *HandlerWrapper) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "CONNECT" {

	} else {
		handler.DumpHTTP(resp, req)
	}
}

// DumpHTTP function to dump the HTTP request header and body
func (handler *HandlerWrapper) DumpHTTP(resp http.ResponseWriter, req *http.Request) {
	req.Header.Del("Proxy-Connection")
	req.Header.Set("Connection", "Keep-Alive")

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		mylog.Fatalln("Fail to dump the http request.")
	}

	connHj, _, err := resp.(http.Hijacker).Hijack()
	if err != nil {
		mylog.Fatalln("Hijack fail to take over the TCP connection from client's request")
	}
	defer connHj.Close()

	host := req.Host
	matched, _ := regexp.MatchString(":[0-9]+$", host)
	if !matched {
		host += ":80"
	}

	connOut, err := net.DialTimeout("tcp", host, time.Second*30)
	if err != nil {
		mylog.Fatalln("Dial to", host, "error:", err)
		return
	}
	// Write writes an HTTP/1.1 request, which is the header and body, in wire format. This method consults the following fields of the request:
	/*
		Host
		URL
		Method (defaults to "GET")
		Header
		ContentLength
		TransferEncoding
		Body
	*/
	if err = req.Write(connOut); err != nil {
		mylog.Fatalln("send to server error", err)
		return
	}

	respFromRemote, err := http.ReadResponse(bufio.NewReader(connOut), req)
	if err != nil && err != io.EOF {
		mylog.Fatalln("Fail to read response from remote server.", err)
	}

	respDump, err := httputil.DumpResponse(respFromRemote, true)
	if err != nil {
		mylog.Fatalln("Fail to dump the response.", err)
	}
	// Send remote response back to client
	_, err = connHj.Write(respDump)
	if err != nil {
		mylog.Fatalln("Fail to send response back to client.", err)
	}

	go httpDump(reqDump, respFromRemote)
}
