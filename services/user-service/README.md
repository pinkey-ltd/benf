# User Service

## gen

```bash
protoc --proto_path="./services/user-service/proto"  --openapi_out="fq_schema_naming=true,default_response=false:./services/user-service/cmd/static" --go_out="./services/user-service/pb" --go-grpc_out="./services/user-service/pb" "./services/user-service/proto/user.proto"

```

## run

```bash
go run .
```