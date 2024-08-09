package logic

import (
	"context"
	"fmt"

	"crud_api/internal/svc"
	"crud_api/internal/types"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCatLogic {
	return &GetCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCatLogic) GetCat(req *types.CatReq) (resp *types.Cat, err error) {
	resp = new(types.Cat)
	cat, e := model.QueryCat(req.Uid)
	resp.Uid = cat.Uid
	resp.Name = cat.Name
	resp.Age = cat.Age
	if e == 500 {
		fmt.Println("查询猫猫信息失败")
		return
	}
	fmt.Println("查询猫猫信息成功")
	return
}
