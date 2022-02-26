package utils

import (
	"errors"
	"io/ioutil"
	"net/http"

	"airport-proxy/logger"
)

func HTTPGet(url string, headers map[string]string) (error, []byte) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil
	}

	for k, v := range headers {
		request.Header.Add(k, v)
	}

	resp, err := client.Do(request)
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	if resp.StatusCode != http.StatusOK {
		logger.Printf("airport response not ok, status code %s, body %s",
			resp.StatusCode, string(respBytes))
		return errors.New("airport response not ok"), nil
	}

	return nil, respBytes
}
