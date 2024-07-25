// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go users
namespace py users
namespace java users
namespace rs users


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Users {

1:i32 id
2:string user_name
3:string email
4:string password
5:string address
6:string private
7:i64 created_at
8:i64 updated_at
9:i64 deleted_at


}

//Create_service
struct CreateUsersRequest{

1:string user_name  (api.body="user_name", api.form="user_name")
2:string email  (api.body="email", api.form="email")
3:string password  (api.body="password", api.form="password")
4:string address  (api.body="address", api.form="address")
5:string private  (api.body="private", api.form="private")


}

struct CreateUsersResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryUsersRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryUsersResponse{
   1: Code code
   2: string msg
   3: list<Users> userss
   4: i32 totoal
}

//Delete_service
struct DeleteUsersRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteUsersResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateUsersRequest{

1:i32 id  (api.body="id", api.form="id")
2:string user_name  (api.body="user_name", api.form="user_name")
3:string email  (api.body="email", api.form="email")
4:string password  (api.body="password", api.form="password")
5:string address  (api.body="address", api.form="address")
6:string private  (api.body="private", api.form="private")


}

struct UpdateUsersResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service UsersService {
   CreateUsersResponse CreateUsers(1:CreateUsersRequest req)(api.post="/api/users/create/")
   QueryUsersResponse  QueryUsers(1: QueryUsersRequest req)(api.get="/api/users/query/")
   DeleteUsersResponse DeleteUsers(1:DeleteUsersRequest req)(api.delete="/api/users/delete/:id")
   UpdateUsersResponse UpdateUsers(1:UpdateUsersRequest req)(api.put="/api/users/update/:id")
}