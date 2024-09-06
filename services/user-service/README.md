# User Service

## gen

```bash
protoc --proto_path="./services/user-service/proto"  --go_out="./services/user-service/pb" --go-grpc_out="./services/user-service/pb" "./services/user-service/proto/user.proto"
```