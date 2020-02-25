package filter

import (
	log "github.com/alecthomas/log4go"
	"github.com/valyala/fasthttp"
	"github.com/wanghongfei/gogate/server"
)

/**
 * 认证过滤器
 */
func AuthMatchPreFilter(s *server.Server, ctx *fasthttp.RequestCtx, newRequest *fasthttp.Request) bool {
	log.Debug("过滤器处理", newRequest)
	return false
}
