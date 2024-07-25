package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	address "towin/hertz_gen/address"
)

type CreateAddressService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateAddressService(Context context.Context, RequestContext *app.RequestContext) *CreateAddressService {
	return &CreateAddressService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateAddressService) Run(req *address.CreateAddressRequest) (resp *address.CreateAddressResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
