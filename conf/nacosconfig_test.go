package conf

import (
	"fmt"
	"github.com/wanghongfei/gogate/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestNacosConfig(t *testing.T) {
	fmt.Println("test nacos config")

	//读取测试的yml数据
	f, err := os.Open("nacosconfig_test.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	fBytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(fBytes))

	//将配置转换成对象
	var nacosConf = new(NacosGatewayConf)
	err = yaml.Unmarshal(fBytes, nacosConf)
	if err != nil {
		fmt.Println(err)
	}
	nacosConf.Init()

	{
		flag := utils.MatchUrl("/jmfen/usercenter/", "/jmfen/usercenter/abc")
		fmt.Println(1, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/11", "/jmfen/usercenter/abc")
		fmt.Println(2, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/*", "/jmfen/usercenter/abc")
		fmt.Println(3, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/*", "/jmfen/usercenter/abc?13df")
		fmt.Println(4, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/*", "/jmfen/usercenter/abc/334")
		fmt.Println(5, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/**", "/jmfen/usercenter/abc/aa334")
		fmt.Println(6, flag)
	}

	{
		flag := utils.MatchUrl("/jmfen/usercenter/**", "/jmfen/usercenter/abc/aa334?1132")
		fmt.Println(7, flag)
	}

	fmt.Println(nacosConf)
}
