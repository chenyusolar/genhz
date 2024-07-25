package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	task "towin/hertz_gen/task"
)

type CreateTaskService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCreateTaskService(Context context.Context, RequestContext *app.RequestContext) *CreateTaskService {
	return &CreateTaskService{RequestContext: RequestContext, Context: Context}
}

func (h *CreateTaskService) Run(req *task.CreateTaskRequest) (resp *task.CreateTaskResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
