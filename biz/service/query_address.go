package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	address "towin/hertz_gen/address"
)

type QueryAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryAddressService(Context context.Context, RequestContext *app.RequestContext) *QueryAddressService {
	return &QueryAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryAddressService) Run(req *address.QueryAddressRequest) (resp *address.QueryAddressResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
