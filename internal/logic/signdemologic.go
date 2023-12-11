package logic

import (
	"context"
	"fmt"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignDemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignDemoLogic {
	return &SignDemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignDemoLogic) SignDemo(req *types.SignDemoReq) (resp *types.SignDemoResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.SignDemoResp)
	resp.Msg = fmt.Sprintf("hello %s", req.Msg)
	return
}
