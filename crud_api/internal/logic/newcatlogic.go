package logic

import (
	"context"
	"fmt"

	"crud_api/internal/svc"
	"crud_api/internal/types"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type NewCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNewCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NewCatLogic {
	return &NewCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NewCatLogic) NewCat(req *types.Cat) (resp *types.ErrResp, err error) {
	resp = new(types.ErrResp)
	var cat model.Cat
	cat.Uid = req.Uid
	cat.Name = req.Name
	cat.Age = req.Age
	resp.Err = model.CreateCat(cat)
	if resp.Err != 200 {
		fmt.Println("新建猫猫信息失败")
		return
	}
	fmt.Println("新建猫猫信息成功")
	return
}
