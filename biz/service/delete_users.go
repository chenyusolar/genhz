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

type DeleteUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteUsersService(Context context.Context, RequestContext *app.RequestContext) *DeleteUsersService {
	return &DeleteUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteUsersService) Run(req *users.DeleteUsersRequest) (resp *users.DeleteUsersResponse, err error) {

	resp = &users.DeleteUsersResponse{
		Code: users.Code_Success,
		Msg:  "",
	}

	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code
	ret := mysql.DB.Delete(&model.User{}, req.ID)
	fmt.Printf("Delete: %+v", ret.Error.Error())

	return resp, ret.Error
}
