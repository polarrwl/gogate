package conf

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/alecthomas/log4go"
	asynclog "github.com/alecthomas/log4go"
	"gopkg.in/yaml.v2"
)

type GateConfig struct {
	Version string `yaml:"version"`

	ServerConfig *ServerConfig `yaml:"server"`
	RedisConfig  *RedisConfig  `yaml:"redis"`

	EurekaConfig *EurekaConfig `yaml:"eureka"`
	ConsulConfig *ConsulConfig `yaml:"consul"`
	NacosConfig  *NacosConfig  `yaml:"nacos"`

	Traffic *TrafficConfig `yaml:"traffic"`
}

type ServerConfig struct {
	AppName       string `yaml:"appName"`
	Host          string `yaml:"host"`
	Port          int    `yaml:"port"`
	MaxConnection int    `yaml:"maxConnection"`
	MaxConnsPerHost	int	`yaml:"maxConnsPerHost"`
	// 请求超时时间, ms
	Timeout int `yaml:"timeout"`
}

type EurekaConfig struct {
	Enable            bool   `yaml:"enable"`
	ConfigFile        string `yaml:"configFile"`
	RouteFile         string `yaml:"routeFile"`
	EvictionDuration  uint   `yaml:"evictionDuration"`
	HeartbeatInterval int    `yaml:"heartbeatInterval"`
}

type ConsulConfig struct {
	Enable  bool   `yaml:"enable"`
	Address string `yaml:"address"`
}

type TrafficConfig struct {
	EnableTrafficRecord bool   `yaml:"enableTrafficRecord"`
	TrafficLogDir       string `yaml:"trafficLogDir"`
}

type NacosConfig struct {
	Enable bool   `yaml:"enable"`
	Server string `yaml:"server"`
	Port   int64  `yaml:"port"`
}

type RedisConfig struct {
	Enabled        bool
	Addr           string
	RateLimiterLua string `yaml:"rateLimiterLua"`
}

var App *GateConfig

func LoadConfig(filename string) {
	f, err := os.Open(filename)
	if nil != err {
		log4go.Error(err)
		panic(err)
	}
	defer f.Close()

	buf, _ := ioutil.ReadAll(f)

	config := new(GateConfig)
	err = yaml.Unmarshal(buf, config)
	if nil != err {
		asynclog.Error(err)
		panic(err)
	}

	validateGogateConfig(config)
}

func InitLog(filename string) {
	log4go.LoadConfiguration(filename)
}

func validateGogateConfig(config *GateConfig) error {
	if nil == config {
		return errors.New("config is nil")
	}

	// 检查eureka配置
	euConfig := config.EurekaConfig
	if nil == euConfig {
		return errors.New("eureka config cannot be empty")
	}
	if euConfig.ConfigFile == "" || euConfig.RouteFile == "" {
		return errors.New("eureka or route config file cannot be empty")
	}

	servCfg := config.ServerConfig
	if servCfg.AppName == "" {
		servCfg.AppName = "gogate"
	}

	if servCfg.Host == "" {
		servCfg.Host = "127.0.0.1"
	}

	if servCfg.Port == 0 {
		servCfg.Port = 8080
	}

	if servCfg.MaxConnection == 0 {
		servCfg.MaxConnection = 1000
	}

	if servCfg.Timeout == 0 {
		servCfg.Timeout = 3000
	}

	trafficCfg := config.Traffic
	if trafficCfg.EnableTrafficRecord {
		if trafficCfg.TrafficLogDir == "" {
			trafficCfg.TrafficLogDir = "/tmp"
		}
	}

	rdConfig := config.RedisConfig
	if rdConfig.Enabled {
		if rdConfig.Addr == "" {
			rdConfig.Addr = "127.0.0.1:6379"
		}
	}

	App = config

	return nil
}
