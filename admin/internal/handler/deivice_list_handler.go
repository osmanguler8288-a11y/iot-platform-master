// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strconv"

	"iot-platform-master/admin/internal/logic"
	"iot-platform-master/admin/internal/svc"
	"iot-platform-master/admin/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeiviceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeiviceListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		req.Page, _ = strconv.Atoi(r.URL.Query().Get("page"))
		req.Size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		req.Name= r.URL.Query().Get("name")
		l := logic.NewDeiviceListLogic(r.Context(), svcCtx)
		resp, err := l.DeiviceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
