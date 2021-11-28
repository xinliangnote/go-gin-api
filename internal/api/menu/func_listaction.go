package menu

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/menu"

	"github.com/spf13/cast"
)

type listActionRequest struct {
	Id string `form:"id"` // hashID
}

type listActionData struct {
	HashId string `json:"hash_id"` // hashID
	MenuId int32  `json:"menu_id"` // 菜单栏ID
	Method string `json:"method"`  // 调用方secret
	API    string `json:"api"`     // 调用方对接人
}

type listActionResponse struct {
	MenuName string           `json:"menu_name"`
	List     []listActionData `json:"list"`
}

// ListAction 功能权限列表
// @Summary 功能权限列表
// @Description 功能权限列表
// @Tags API.menu
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id query string true "hashID"
// @Success 200 {object} listActionResponse
// @Failure 400 {object} code.Failure
// @Router /api/menu_action [get]
// @Security LoginToken
func (h *handler) ListAction() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listActionRequest)
		res := new(listActionResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithError(err),
			)
			return
		}

		id := int32(ids[0])

		searchOneData := new(menu.SearchOneData)
		searchOneData.Id = id

		menuInfo, err := h.menuService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuDetailError,
				code.Text(code.MenuDetailError)).WithError(err),
			)
			return
		}

		res.MenuName = menuInfo.Name

		searchListData := new(menu.SearchListActionData)
		searchListData.MenuId = menuInfo.Id

		resListData, err := h.menuService.ListAction(c, searchListData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.AuthorizedListAPIError,
				code.Text(code.AuthorizedListAPIError)).WithError(err),
			)
			return
		}

		res.List = make([]listActionData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithError(err),
				)
				return
			}

			data := listActionData{
				HashId: hashId,
				MenuId: v.MenuId,
				Method: v.Method,
				API:    v.Api,
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
