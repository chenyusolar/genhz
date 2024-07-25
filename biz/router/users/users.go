// Code generated by hertz generator. DO NOT EDIT.

package users

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	users "towin/biz/handler/users"
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
			_users := _api.Group("/users", _usersMw()...)
			{
				_create := _users.Group("/create", _createMw()...)
				_create.POST("/", append(_createusersMw(), users.CreateUsers)...)
			}
			{
				_delete := _users.Group("/delete", _deleteMw()...)
				_delete.DELETE("/:id", append(_deleteusersMw(), users.DeleteUsers)...)
			}
			{
				_query := _users.Group("/query", _queryMw()...)
				_query.GET("/", append(_queryusersMw(), users.QueryUsers)...)
			}
			{
				_update := _users.Group("/update", _updateMw()...)
				_update.PUT("/:id", append(_updateusersMw(), users.UpdateUsers)...)
			}
		}
	}
}
