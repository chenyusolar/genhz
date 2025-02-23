// Code generated by hertz generator. DO NOT EDIT.

package task

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	task "towin/biz/handler/task"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		{
			_task := _api.Group("/task", _taskMw()...)
			{
				_create := _task.Group("/create", _createMw()...)
				_create.POST("/", append(_createtaskMw(), task.CreateTask)...)
			}
			{
				_delete := _task.Group("/delete", _deleteMw()...)
				_delete.DELETE("/:id", append(_deletetaskMw(), task.DeleteTask)...)
			}
			{
				_query := _task.Group("/query", _queryMw()...)
				_query.GET("/", append(_querytaskMw(), task.QueryTask)...)
			}
			{
				_update := _task.Group("/update", _updateMw()...)
				_update.PUT("/:id", append(_updatetaskMw(), task.UpdateTask)...)
			}
		}
	}
}
