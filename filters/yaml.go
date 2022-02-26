package filters

import (
	"gopkg.in/yaml.v2"

	"cnf/common"
	"cnf/logger"
)

// YamlFilter Yaml类型的过滤，针对Yaml格式，格式见demo.yaml
type YamlFilter struct {
	Content string
}

func (r YamlFilter) Do(keywords []string) (error, []byte) {
	conf := &common.YamlConf{}
	if err := yaml.Unmarshal([]byte(r.Content), conf); err != nil {
		logger.Printf("fail to parse yaml %s, error %s", r.Content, err.Error())
		return err, nil
	}

	newConf := remove(*conf, keywords)
	strBytes, err := yaml.Marshal(newConf)
	if err != nil {
		logger.Printf("fail to marshal yaml, error %s", err.Error())
		return err, nil
	}

	return nil, strBytes
}

func remove(conf common.YamlConf, keywords []string) common.YamlConf {
	newConf := conf
	var proxies []common.Proxy
	var proxyGroups []common.ProxyGroup

	for _, p := range conf.Proxies {
		if containsKeywords(p.Name, keywords) {
			continue
		}
		proxies = append(proxies, p)
	}
	for _, group := range conf.ProxyGroups {
		newGroup := group
		var groupProxies []string
		for _, p := range group.Proxies {
			if containsKeywords(p, keywords) {
				continue
			}
			groupProxies = append(groupProxies, p)
		}
		newGroup.Proxies = groupProxies
		proxyGroups = append(proxyGroups, newGroup)
	}

	newConf.Proxies = proxies
	newConf.ProxyGroups = proxyGroups
	return newConf
}
