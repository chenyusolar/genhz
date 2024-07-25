package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	task "towin/hertz_gen/task"
)

type DeleteTaskService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewDeleteTaskService(Context context.Context, RequestContext *app.RequestContext) *DeleteTaskService {
	return &DeleteTaskService{RequestContext: RequestContext, Context: Context}
}

func (h *DeleteTaskService) Run(req *task.DeleteTaskRequest) (resp *task.DeleteTaskResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
