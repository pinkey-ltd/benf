# BenF
犇 project gofr

## dependents

```bash

go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
# go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

## Run

``` bash
go run ./cmd
```

## todo
 [ ] Cesium Ion Access Token
 [ ] Default Views Widget Controls 
 
## Step
```mermaid
graph TB
    ProtoFile(编写proto文件) --> Generate[Makefile]
    Generate -->|生成| PB[grpc接口]
    Generate -->|生成| OpenAPI[swagger]
    Generate --> Login[gRPC-Gateway]
    OpenAPI --> |生成| APIHandler[REST接口]

 ```

```mermaid
sequenceDiagram
    participant User
    participant OLTP
    User->>John: Hello John, how are you?
    loop HealthCheck
John->>John: Fight against hypochondria
end
Note right of John: Rational thoughts <br/>prevail!
John-->>Alice: Great!
John->>Bob: How about you?
Bob-->>John: Jolly good!

```