package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	address "towin/hertz_gen/address"
)

type DeleteAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteAddressService(Context context.Context, RequestContext *app.RequestContext) *DeleteAddressService {
	return &DeleteAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteAddressService) Run(req *address.DeleteAddressRequest) (resp *address.DeleteAddressResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
