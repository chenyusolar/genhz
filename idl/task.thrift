// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go task
namespace py task
namespace java task
namespace rs task


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Task {

1:i32 id
2:i32 userid
3:string adress
4:string token
5:string recipients
6:string amounts
7:i8 status
8:i64 created_at
9:i64 updated_at
10:i64 deleted_at


}

//Create_service
struct CreateTaskRequest{

1:i32 userid  (api.body="userid", api.form="userid")
2:string adress  (api.body="adress", api.form="adress")
3:string token  (api.body="token", api.form="token")
4:string recipients  (api.body="recipients", api.form="recipients")
5:string amounts  (api.body="amounts", api.form="amounts")
6:i8 status  (api.body="status", api.form="status")


}

struct CreateTaskResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryTaskRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryTaskResponse{
   1: Code code
   2: string msg
   3: list<Task> tasks
   4: i32 totoal
}

//Delete_service
struct DeleteTaskRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteTaskResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateTaskRequest{

1:i32 id  (api.body="id", api.form="id")
2:i32 userid  (api.body="userid", api.form="userid")
3:string adress  (api.body="adress", api.form="adress")
4:string token  (api.body="token", api.form="token")
5:string recipients  (api.body="recipients", api.form="recipients")
6:string amounts  (api.body="amounts", api.form="amounts")
7:i8 status  (api.body="status", api.form="status")


}

struct UpdateTaskResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service TaskService {
   CreateTaskResponse CreateTask(1:CreateTaskRequest req)(api.post="/api/task/create/")
   QueryTaskResponse  QueryTask(1: QueryTaskRequest req)(api.get="/api/task/query/")
   DeleteTaskResponse DeleteTask(1:DeleteTaskRequest req)(api.delete="/api/task/delete/:id")
   UpdateTaskResponse UpdateTask(1:UpdateTaskRequest req)(api.put="/api/task/update/:id")
}