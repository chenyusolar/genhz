package users

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"towin/biz/service"
	"towin/biz/utils"
	users "towin/hertz_gen/users"
)

// CreateUsers .
// @router /api/users/create/ [POST]
func CreateUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.CreateUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewCreateUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// QueryUsers .
// @router /api/users/query/ [GET]
func QueryUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.QueryUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Printf("Query:  %+v", &req)
	resp, err := service.NewQueryUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// DeleteUsers .
// @router /api/users/delete/:id [DELETE]
func DeleteUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.DeleteUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Printf("Delete:  %+v", &req)
	resp, err := service.NewDeleteUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateUsers .
// @router /api/users/update/:id [PUT]
func UpdateUsers(ctx context.Context, c *app.RequestContext) {
	var err error
	var req users.UpdateUsersRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	fmt.Printf("Update:  %+v", &req)
	resp, err := service.NewUpdateUsersService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
