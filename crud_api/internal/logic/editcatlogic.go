package logic

import (
	"context"
	"fmt"

	"crud_api/internal/svc"
	"crud_api/internal/types"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditCatLogic {
	return &EditCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditCatLogic) EditCat(req *types.Cat) (resp *types.ErrResp, err error) {
	resp = new(types.ErrResp)
	var cat model.Cat
	cat.Uid = req.Uid
	cat.Name = req.Name
	cat.Age = req.Age
	resp.Err = model.EditCat(cat)
	if resp.Err != 200 {
		fmt.Println("修改猫猫信息失败")
		return
	}
	fmt.Println("修改猫猫信息成功")
	return
}
