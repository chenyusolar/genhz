package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	task "towin/hertz_gen/task"
)

type QueryTaskService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewQueryTaskService(Context context.Context, RequestContext *app.RequestContext) *QueryTaskService {
	return &QueryTaskService{RequestContext: RequestContext, Context: Context}
}

func (h *QueryTaskService) Run(req *task.QueryTaskRequest) (resp *task.QueryTaskResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
