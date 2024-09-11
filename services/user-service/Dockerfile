FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go env -w GOPROXY="https://goproxy.cn,direct" \
    && go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd

# Final stage
FROM alpine:latest

WORKDIR /root/

FROM alpine:latest

RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

# configs
COPY --from=builder /app/cmd/configs .

COPY --from=builder /app/main .

EXPOSE 8000-8001
EXPOSE 2121

CMD ["./main"]