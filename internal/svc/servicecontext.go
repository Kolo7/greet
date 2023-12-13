package svc

import (
	"greet/internal/config"
	"greet/internal/middleware"
	"greet/model/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	UserAgentMiddleware rest.Middleware
	MjHistorysModel     mysql.MjHistorysModel
	sqlConn             sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DBConn.Dsn)

	return &ServiceContext{
		Config:              c,
		UserAgentMiddleware: middleware.NewUserAgentMiddleware().Handle,
		MjHistorysModel:     mysql.NewMjHistorysModel(sqlConn),
		sqlConn:             sqlConn,
	}
}
