package handler

import (
	"net/http"

	"backend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

// 处理预检请求
func Prefix_Managing(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if r.Method == "OPTIONS" || r.Method == "POST" || r.Method == "GET" {
		// 这是预检请求 ，在POST请求也对其进行处理
		(w).Header().Set("Access-Control-Allow-Origin", origin) // 允许指定来源的跨域请求
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		(w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-Extra-Header, Content-Type, Accept, Authorization")
		(w).Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		(w).Header().Set("Access-Control-Allow-Credentials", "true")
	}

}
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/filelist/:name",
				Handler: BackendHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/filelist/:name",
				Handler: Prefix_Managing,
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/file",
				Handler: BackendHandlerPOST(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/file",
				Handler: Prefix_Managing,
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/save",
				Handler: BackendHandlerSAVE(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/save",
				Handler: Prefix_Managing,
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/export",
				Handler: BackendHandlerEXPORT(serverCtx),
			},
		},
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodOptions,
				Path:    "/export",
				Handler: Prefix_Managing,
			},
		},
	)
}
