package conf

import "strings"

/**
 * 网关nacos配置
 */
type NacosGatewayConf struct {
	Zlt *NacosGatewayZlt `yaml:"zlt"`
}

type NacosGatewayZlt struct {
	Security *NacosGatewaySecurity `yaml:"security"`
}

type NacosGatewaySecurity struct {
	Ignore *NacosGatewayUrl    `yaml:"ignore"`
	Forbid *NacosGatewayUrl    `yaml:"forbid"`
	System *NacosGatewaySystem `yaml:"system"`
	Auth   *NacosGatewayUrl    `yaml:"auth"`
}

type NacosGatewaySystem struct {
	Ignore *NacosGatewayUrl `yaml:"ignore"`
}

type NacosGatewayUrl struct {
	HttpUrlStr string `yaml:"httpUrls"`
	HttpUrls   []string
	UrlEnabled bool `yaml:"urlEnabled"`
}

func (nacosConf *NacosGatewayConf) Init() {
	//将str解析成array
	if nacosConf.Zlt != nil && nacosConf.Zlt.Security != nil {
		nacosConf.Zlt.Security.Ignore.analyseUrl()
		nacosConf.Zlt.Security.Auth.analyseUrl()
		nacosConf.Zlt.Security.Forbid.analyseUrl()
		nacosConf.Zlt.Security.System.Ignore.analyseUrl()
	}
}

func (confUrl *NacosGatewayUrl) analyseUrl() {
	stringArr := strings.Split(confUrl.HttpUrlStr, ",")
	for _, v := range stringArr {
		confUrl.HttpUrls = append(confUrl.HttpUrls, strings.TrimSpace(v))
	}
}
