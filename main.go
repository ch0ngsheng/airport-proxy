package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"cnf/common"
	_ "cnf/platforms"
)

var port string

func main() {
	if err := common.ParseConf(); err != nil {
		os.Exit(1)
	}

	flag.StringVar(&port, "port", "8090", "HTTP port to listen")
	fmt.Printf("service is listening on port %s\n", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), common.RequestHandler{})
}
