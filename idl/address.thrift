// refer to  https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/toolkit/

namespace go address
namespace py address
namespace java address
namespace rs address


enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}


//table items
struct Address {

1:i32 id
2:string address
3:string address_type
4:i32 address_owner
5:string private
6:i64 created_at
7:i64 updated_at
8:i64 deleted_at


}

//Create_service
struct CreateAddressRequest{

1:string address  (api.body="address", api.form="address")
2:string address_type  (api.body="address_type", api.form="address_type")
3:i32 address_owner  (api.body="address_owner", api.form="address_owner")
4:string private  (api.body="private", api.form="private")


}

struct CreateAddressResponse{
   1: Code code
   2: string msg
}

//Query_service
struct QueryAddressRequest{
   1: optional string Keyword (api.body="keyword", api.form="keyword",api.query="keyword")
   2: i32 page (api.body="page", api.form="page",api.query="page",api.vd="$ > 0")
   3: i32 page_size (api.body="page_size", api.form="page_size",api.query="page_size",api.vd="($ > 0 || $ <= 100)")
}

struct QueryAddressResponse{
   1: Code code
   2: string msg
   3: list<Address> addresss
   4: i32 totoal
}

//Delete_service
struct DeleteAddressRequest{
   1: i32    id   (api.path="id",api.vd="$>0")
}

struct DeleteAddressResponse{
   1: Code code
   2: string msg
}

//Update_service
struct UpdateAddressRequest{

1:i32 id  (api.body="id", api.form="id")
2:string address  (api.body="address", api.form="address")
3:string address_type  (api.body="address_type", api.form="address_type")
4:i32 address_owner  (api.body="address_owner", api.form="address_owner")
5:string private  (api.body="private", api.form="private")


}

struct UpdateAddressResponse{
   1: Code code
   2: string msg
}

//Define Service Routine
service AddressService {
   CreateAddressResponse CreateAddress(1:CreateAddressRequest req)(api.post="/api/address/create/")
   QueryAddressResponse  QueryAddress(1: QueryAddressRequest req)(api.get="/api/address/query/")
   DeleteAddressResponse DeleteAddress(1:DeleteAddressRequest req)(api.delete="/api/address/delete/:id")
   UpdateAddressResponse UpdateAddress(1:UpdateAddressRequest req)(api.put="/api/address/update/:id")
}