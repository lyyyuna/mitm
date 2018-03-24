package mitm

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func httpDump(reqDump []byte, resp *http.Response) {
	defer resp.Body.Close()
	respStatus := resp.Status

	req, _ := ParseReq(reqDump)

	fmt.Println("Status: ", respStatus)
	fmt.Printf("%s %s %s\n", req.Method, req.Host, req.RequestURI)
	fmt.Printf("%s %s\n", "Remote Addr: ", req.RemoteAddr)
	for headerName, headerContext := range req.Header {
		fmt.Printf("\t%s: %s\n", headerName, headerContext)
	}

	if req.Method == "POST" {
		fmt.Println("POST Param:")
		err := req.ParseForm()
		if err != nil {
			logger.Println("ParseForm error: ", err)
		} else {
			for k, v := range req.Form {
				fmt.Printf("\t\t%s: %s\n", k, v)
			}
		}
	}

	fmt.Println("Response:")
	for headerName, headerContext := range resp.Header {
		fmt.Printf("\t%s: %s\n", headerName, headerContext)
	}
	fmt.Println("")
}

// ParseReq why
func ParseReq(b []byte) (*http.Request, error) {
	// func ReadRequest(b *bufio.Reader) (req *Request, err error) { return readRequest(b, deleteHostHeader) }
	var buf io.ReadWriter
	buf = new(bytes.Buffer)
	buf.Write(b)
	bufr := bufio.NewReader(buf)
	return http.ReadRequest(bufr)
}
