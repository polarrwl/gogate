package discovery

import (
	"errors"
	alertlog "github.com/alecthomas/log4go"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/wanghongfei/gogate/conf"
	"log"
	"net"
	"os"
)

var namingClientObj naming_client.INamingClient

func InitNacosClient() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	alertlog.Info("初始化nacos client")

	// 可以没有，采用默认值
	clientConfig := constant.ClientConfig{
		TimeoutMs:            10 * 1000,
		ListenInterval:       30 * 1000,
		BeatInterval:         5 * 1000,
		NamespaceId:          "public", //nacos命名空间
		UpdateCacheWhenEmpty: true,     //空的也进行更新，不然会有些服务空的会不停的打印日志
		LogDir:               "./logs",
		CacheDir:             "./cache",
	}

	// 至少一个
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      conf.App.NacosConfig.Server,
			ContextPath: "/nacos",
			Port:        uint64(conf.App.NacosConfig.Port),
		},
	}

	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		alertlog.Error("nacos注册失败", err)
		panic(err)
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		alertlog.Error("nacos配置失败", err)
		panic(err)
	}

	//nacos的日志输出有问题，这里需要强行扭转过来
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)

	//nacos的日志输出有问题，这里需要强行扭转过来
	alertlog.Info("nacos client创建:", namingClient, configClient)

	clientIp, _ := getClientIp()

	//调用nacso注册实例（这里会自动启动心跳检查线程）
	success, _ := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          clientIp,
		Port:        uint64(conf.App.ServerConfig.Port),
		ServiceName: conf.App.ServerConfig.AppName,
		Weight:      10,
		//ClusterName: "a",
		Enable:    true,
		Healthy:   true,
		Ephemeral: true,
	})

	alertlog.Info("nacos客戶端注冊結果:", success)

	namingClientObj = namingClient
}

/**
获取本机IP地址
*/
func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("Can not find the client ip address!")

}
