package mitm

import (
	"fmt"
	"mylog"
	"net/http"
)

func httpDump(reqDump []byte, req *http.Request, resp *http.Response) {
	defer resp.Body.Close()
	respStatus := resp.Status

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
			mylog.Fatalln("ParseForm error: ", err)
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
