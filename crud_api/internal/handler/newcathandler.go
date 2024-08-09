package handler

import (
	"net/http"

	"crud_api/internal/logic"
	"crud_api/internal/svc"
	"crud_api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NewCatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Cat
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNewCatLogic(r.Context(), svcCtx)
		resp, err := l.NewCat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
