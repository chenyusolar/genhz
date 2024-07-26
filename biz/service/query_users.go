package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"towin/biz/dal/model"
	"towin/biz/dal/mysql"

	users "towin/hertz_gen/users"
)

type QueryUsersService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryUsersService(Context context.Context, RequestContext *app.RequestContext) *QueryUsersService {
	return &QueryUsersService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryUsersService) Run(req *users.QueryUsersRequest) (resp *users.QueryUsersResponse, err error) {

	resp = &users.QueryUsersResponse{}

	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	// todo edit your code

	fmt.Println("req=", req)
	fmt.Println("resp=", resp)

	ret := mysql.DB.Find(&[]model.User{})

	if ret.RowsAffected > 0 {
		resp.Code = users.Code_Success
		resp.Msg = "success:"
		fmt.Printf("user list: %+v", ret)
	} else {
		resp.Code = users.Code_DBErr
		resp.Msg = "DBErr:" + ret.Error.Error()
	}

	return resp, ret.Error
}
