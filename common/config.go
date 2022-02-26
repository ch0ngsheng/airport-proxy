package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	jsonFile = "config.json"
)

var AirPorts = map[string]*AirPort{}

type AirPort struct {
	Name     string            `json:"name"`
	Url      string            `json:"url"`
	Method   string            `json:"method"`
	Params   []*AirPortParamKV `json:"params"`
	Keywords []string          `json:"keywords"`
}

type AirPortParamKV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ParseConf() error {
	bytes, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("需要放置config.json文件", err)
		return err
	}

	type AirPortObj struct {
		APs []*AirPort `json:"airports"`
	}

	aps := &AirPortObj{}
	err = json.Unmarshal(bytes, aps)
	if err != nil {
		fmt.Println("config.json格式有误", err)
		return err
	}

	for _, ap := range aps.APs {
		AirPorts[strings.ToLower(ap.Name)] = ap
	}
	return nil
}
