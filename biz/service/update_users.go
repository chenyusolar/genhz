package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"towin/biz/dal/model"
	"towin/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
	users "towin/hertz_gen/users"
)

type UpdateUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUsersService(Context context.Context, RequestContext *app.RequestContext) *UpdateUsersService {
	return &UpdateUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUsersService) Run(req *users.UpdateUsersRequest) (resp *users.UpdateUsersResponse, err error) {
	resp = &users.UpdateUsersResponse{
		Code: users.Code_Success,
		Msg:  "",
	}

	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// todo edit your code
	ret := mysql.DB.Model(&model.User{}).Where("id", req.ID).Updates(&model.User{UserName: req.UserName, Email: req.Email, Password: req.Password, Address: req.Address, Private: req.Private})
	if ret.RowsAffected > 0 {
		resp.Code = users.Code_Success
		resp.Msg = "success"
		fmt.Println("Update Succeed:", ret.Error.Error())
	} else {
		resp.Code = users.Code_DBErr
		resp.Msg = "Failed"
		fmt.Println("Update Failed:", ret.Error.Error())

	}

	return resp, ret.Error
}
