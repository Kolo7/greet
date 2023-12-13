package user

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet(req *types.Request) (resp *types.Response, err error) {
	if req.Name == "you" {
		return nil, errors.New(400, "搞错了")
	}
	mjHistory, err := l.svcCtx.MjHistorysModel.FindOne(l.ctx, 14201)
	if err != nil {
		errors.New(10000, "数据库有问题")
	}
	return &types.Response{Message: mjHistory.Prompt}, nil
}
