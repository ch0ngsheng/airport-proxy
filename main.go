package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"airport-proxy/common"
	_ "airport-proxy/platforms"
)

var port string

func main() {
	if err := common.ParseConf(); err != nil {
		os.Exit(1)
	}

	flag.StringVar(&port, "port", "9090", "HTTP port to listen")
	fmt.Printf("service is listening on port %s\n", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), common.RequestHandler{})
}
