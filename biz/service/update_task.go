package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	task "towin/hertz_gen/task"
)

type UpdateTaskService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateTaskService(Context context.Context, RequestContext *app.RequestContext) *UpdateTaskService {
	return &UpdateTaskService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateTaskService) Run(req *task.UpdateTaskRequest) (resp *task.UpdateTaskResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
