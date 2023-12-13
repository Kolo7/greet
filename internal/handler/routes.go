// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "greet/internal/handler/user"
	userInfo "greet/internal/handler/userInfo"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAgentMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/from/:name",
					Handler: user.GreetHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sign/demo",
				Handler: SignDemoHandler(serverCtx),
			},
		},
		rest.WithSignature(serverCtx.Config.Signature),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserAgentMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: userInfo.UserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: pingHandler(serverCtx),
			},
		},
	)
}
