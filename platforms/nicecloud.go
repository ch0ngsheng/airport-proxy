package platforms

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"airport-proxy/common"
	"airport-proxy/filters"
	"airport-proxy/logger"
	"airport-proxy/utils"
)

func init() {
	common.HandlerMap[strings.ToLower(NC{}.Name())] = &NC{}
}

// NC NiceCloud
type NC struct {
}

func (N NC) RemoveRateNode() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Printf("NiceCloud job...")

		err, respStr := N.reqAirport(request)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		var f filters.Filter
		if fromClash(request.Header.Get("User-Agent")) {
			// 机场会判断，如果请求来自Clash客户端，则返回yaml格式的节点数据
			f = &filters.YamlFilter{Content: respStr}
		} else {
			f = &filters.RawFilter{Content: respStr}
		}
		err, respByte := f.Do(common.AirPorts[strings.ToLower(N.Name())].Keywords)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(respByte)

		logger.Printf("NiceCloud done")
	}
}

func (N NC) Name() string {
	return "NiceCloud"
}

func (N NC) reqAirport(request *http.Request) (error, string) {
	ap := common.AirPorts[strings.ToLower(N.Name())]
	if ap == nil {
		logger.Printf("%s机场的name配置项有误\n", N.Name())
		return errors.New(""), ""
	}

	url := ap.Url + "?"
	for _, p := range ap.Params {
		if p.Key == "token" {
			p.Value = request.URL.Query().Get("token")
		}
		url = fmt.Sprintf("%s%s=%s&", url, p.Key, p.Value)
	}

	headers := map[string]string{}
	headers["User-Agent"] = request.Header.Get("User-Agent")

	err, respBytes := utils.HTTPGet(url, headers)
	if err != nil {
		logger.Printf("fail to request %s, error %s", N.Name(), err.Error())
		return err, ""
	}
	return nil, string(respBytes)
}
