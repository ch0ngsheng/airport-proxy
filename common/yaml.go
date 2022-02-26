package common

// YamlConf 机场返回的yaml格式结构体
type YamlConf struct {
	MixedPort    int          `yaml:"mixed-port"`
	RedirPort    int          `yaml:"redir-port"`
	BindAddress  string       `yaml:"bind-address"`
	AllowLan     bool         `yaml:"allow-lan"`
	Mode         string       `yaml:"mode"`
	LogLevel     string       `yaml:"log-level"`
	ExternalCtrl string       `yaml:"external-controller"`
	Secret       string       `yaml:"secret"`
	Proxies      []Proxy      `yaml:"proxies"`
	ProxyGroups  []ProxyGroup `yaml:"proxy-groups"`
	Dns          DNS          `yaml:"dns"`
	Rules        []string     `yaml:"rules"`
}

type Proxy struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	Udp      bool   `yaml:"udp"`
	SNI      string `yaml:"sni"`
	UUID     string `yaml:"uuid"`
	AlterID  string `yaml:"alterId"`
	Cipher   string `yaml:"cipher"`
	Network  string `yaml:"network"`
}
type ProxyGroup struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Proxies  []string `yaml:"proxies"`
	Url      string   `yaml:"url"`
	Interval int      `yaml:"interval"`
}
type FallbackFilter struct {
	GeoIP  bool     `yaml:"geoip"`
	ipcidr []string `yaml:"ipcidr"`
}
type DNS struct {
	Enable            bool           `yaml:"enable"`
	Ipv6              bool           `yaml:"ipv6"`
	Listen            string         `yaml:"listen"`
	EnhancedMode      string         `yaml:"enhanced-mode"`
	DefaultNameServer []string       `yaml:"default-nameserver"`
	NameServer        []string       `yaml:"nameserver"`
	FakeIPRange       string         `yaml:"fake-ip-range"`
	UseHosts          bool           `yaml:"use-hosts"`
	Fallback          []string       `yaml:"fallback"`
	FbFilter          FallbackFilter `yaml:"fallback-filter"`
}
