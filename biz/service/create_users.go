package service

import (
	"context"
	"towin/biz/dal/model"
	"towin/biz/dal/mysql"
	users "towin/hertz_gen/users"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CreateUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateUsersService(Context context.Context, RequestContext *app.RequestContext) *CreateUsersService {
	return &CreateUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateUsersService) Run(req *users.CreateUsersRequest) (resp *users.CreateUsersResponse, err error) {

	resp = &users.CreateUsersResponse{
		Code: users.Code_Success,
		Msg:  "",
	}

	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	ret := mysql.DB.Create(&model.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
		Address:  req.Address,
		Private:  req.Private,
	})

	if ret.RowsAffected > 0 {
		resp.Code = users.Code_Success
		resp.Msg = "user insert success: "
	} else {
		resp.Code = users.Code_DBErr
		resp.Msg = "user insert fail:"
	}

	return resp, ret.Error
}
