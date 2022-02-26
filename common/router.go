package common

import (
	"net/http"
	"strings"
)

const (
	filterURI = "/v1/ap/filter"

	airportQueryKey = "ap"
)

type RequestHandler struct {
}

func (A RequestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != filterURI && request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	apName := request.URL.Query().Get(airportQueryKey)
	if h, ok := HandlerMap[strings.ToLower(apName)]; ok {
		h.RemoveRateNode()(writer, request)
		return
	}

	writer.WriteHeader(http.StatusNotFound)
}
