package logic

import (
	"context"
	"fmt"

	"crud_api/internal/svc"
	"crud_api/internal/types"
	"crud_api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCatLogic {
	return &DeleteCatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCatLogic) DeleteCat(req *types.CatReqPost) (resp *types.ErrResp, err error) {
	resp = new(types.ErrResp)
	resp.Err = model.DeleteCat(req.Uid)
	if resp.Err != 200 {
		fmt.Println("删除猫猫信息失败")
		return
	}
	fmt.Println("删除猫猫信息成功")
	return
}
