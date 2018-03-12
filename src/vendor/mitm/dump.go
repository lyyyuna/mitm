package mitm

import (
	"fmt"
	"net/http"
)

func httpDump(reqDump []byte, req *http.Request, resp *http.Response) {
	defer resp.Body.Close()
	respStatus := resp.Status

	fmt.Println("Status: ", respStatus)
	fmt.Printf("%s %s\n", req.Method, req.Host+req.RequestURI)
}
