package filters

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/url"
	"strings"

	"airport-proxy/logger"
)

const (
	protocolVmess = iota
	protocolSs
	protocolTrojan
)

func init() {
	verifierMap[protocolVmess] = verifyVmess
	verifierMap[protocolSs] = verifySs
	verifierMap[protocolTrojan] = verifyTrojan
}

type verifier func(str string, keywords []string) (error, bool)

var (
	verifierMap = map[int]verifier{}
)

// RawFilter Raw类型的过滤，针对每行一个节点的格式，格式见raw.txt
type RawFilter struct {
	Content string
}

func (r RawFilter) Do(keywords []string) (error, []byte) {
	decoded, err := base64.StdEncoding.DecodeString(r.Content)
	if err != nil {
		logger.Printf("fail to base64 decode raw response %s, error %s", r.Content, err.Error())
		return err, nil
	}

	err, str := filter(string(decoded), keywords)
	if err != nil {
		return err, nil
	}

	return nil, []byte(base64.StdEncoding.EncodeToString([]byte(str)))
}

func filter(content string, keywords []string) (error, string) {
	var sb strings.Builder

	list := strings.Split(content, "\n")
	for _, node := range list {
		if len(node) == 0 {
			continue
		}
		err, ok := verifierMap[getProtocolType(node)](node, keywords)
		if err != nil {
			logger.Printf("failed to parse %s", node)
			return err, ""
		}
		if !ok {
			continue
		}
		sb.Write([]byte(node))
		sb.Write([]byte("\n"))
	}
	return nil, sb.String()
}

func getProtocolType(str string) int {
	if strings.HasPrefix(str, "vmess") {
		return protocolVmess
	} else if strings.HasPrefix(str, "ss") {
		return protocolSs
	}
	return protocolTrojan
}

// verifyVmess 判断vmess协议的一个节点是否应被过滤
func verifyVmess(str string, keywords []string) (error, bool) {
	// vmess://eyJ2IjoiMiIsInBzIjoiXHVkODNjXHVkZGVmXHVkODNjXHVkZGY1IFx1NjVlNVx1NjcyY1x1NmUzOFx1NjIwZlx1N2ViZlx1Nzk4MVx1ODljNlx1OTg5MXg1MFx1NTAwZFx1NzM4NyoiLCJhZGQiOiJzaC1qcC1pbi5pcGxjLmlua2l3aWtlLmNvbSIsInBvcnQiOiIzMjUzNSIsImlkIjoiMzkzNDFhNTItYjUzYy00OGE4LTgwMzMtYTY5MzY1NjU0M2NhIiwiYWlkIjoiMCIsIm5ldCI6InRjcCIsInR5cGUiOiJub25lIiwiaG9zdCI6IiIsInBhdGgiOiIiLCJ0bHMiOiIifQ==
	base64Encoded := strings.TrimPrefix(str, "vmess://")
	jsonBytes, err := base64.StdEncoding.DecodeString(base64Encoded)
	if err != nil {
		logger.Printf("fail to decode base64 string %s, error %s", base64Encoded, err.Error())
		return err, false
	}

	// json格式为
	// {"v":"2","ps":"\ud83c\udded\ud83c\uddf0\u9999\u6e2f\u6e38\u620f\u7ebf\u7981\u89c6\u9891x50\u500d\u7387*","add":"sz-hk-in.iplc.inkiwike.com","port":"32652","id":"39341a52-b53c-48a8-8033-a693656543ca","aid":"0","net":"tcp","type":"none","host":"","path":"","tls":""}
	type O struct {
		Ps string `json:"ps"`
	}

	var oo = &O{}
	err = json.Unmarshal(jsonBytes, oo)
	if err != nil {
		logger.Printf("fail to unmarshl %s to json when parse vmess of %s", jsonBytes)
		return err, false
	}
	return nil, !containsKeywords(oo.Ps, keywords)
}

// verifySs 判断ss协议的一个节点是否应被过滤
func verifySs(str string, keywords []string) (error, bool) {
	// ss://YWVzLTEyOC1nY206MzkzNDFhNTItYjUzYy00OGE4LTgwMzMtYTY5MzY1NjU0M2Nh@120.240.86.116:45101#%F0%9F%87%BA%F0%9F%87%B8%20%E7%BE%8E%E5%9B%BD-%E6%B3%B8%E7%BE%8E
	urlEncodeds := strings.Split(str, "#")
	if len(urlEncodeds) != 2 {
		logger.Printf("fail to parse ss protocol string %s", str)
		return errors.New("error parse ss"), false
	}
	urlEncoded := urlEncodeds[1]
	decodedBytes, err := url.QueryUnescape(urlEncoded)
	if err != nil {
		logger.Printf("fail to url decode ss protocol string %s, error %s", urlEncoded, err.Error())
		return err, false
	}

	return nil, !containsKeywords(decodedBytes, keywords)
}

// verifyTrojan 判断trojan协议的一个节点是否应被过滤
func verifyTrojan(str string, keywords []string) (error, bool) {
	// trojan://39341a52-b53c-48a8-8033-a693656543ca@1.1.1.1.1:11?allowInsecure=0#%E8%AF%B7%E6%AF%8F%E5%A4%A9%E6%9B%B4%E6%96%B0%E8%AE%A2%E9%98%85%E4%B8%80%E6%AC%A1
	return verifySs(str, keywords)
}
