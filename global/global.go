package global

import (
	log "github.com/alecthomas/log4go"
	serv "github.com/wanghongfei/gogate/server"
	"github.com/wanghongfei/gogate/server/route"
)

//主要是为了防止循环依赖
type QueryGlobal struct {
	Server *serv.Server
}

func (queryGlobal QueryGlobal) AddServerInfo(name string) {
	router := queryGlobal.Server.Router
	if router.Match("/"+name) == nil {
		info := &route.ServiceInfo{
			Id:          name,
			Prefix:      "/" + name,
			StripPrefix: true,
			Name:        name,
		}
		log.Info("添加服务路由信息", info)
		router.AddServerInfo(info)
	} else {
		log.Info("已存在的路由，不用添加", name)
	}
}
