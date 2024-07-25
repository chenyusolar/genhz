package service

import (
	"context"

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

	err := mysql.DB.CreateUsers(h.Context, req)

	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	return
}
