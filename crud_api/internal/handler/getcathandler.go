package handler

import (
	"net/http"

	"crud_api/internal/logic"
	"crud_api/internal/svc"
	"crud_api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCatLogic(r.Context(), svcCtx)
		resp, err := l.GetCat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
