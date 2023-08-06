FROM golang:1.20-alpine AS builder

# RUN apt update && apt install -y  upx
RUN apk add upx

WORKDIR "/go/src/github.com/byeol-i/battery-level-checker/"

COPY go.mod .
COPY go.sum .

RUN mkdir -p /opt/bin/app/

# COPY vendor/ /go/src/

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/app/main cmd/cache/main.go \    
#     && upx --lzma /opt/bin/app/main

RUN go mod tidy \
    && go get -u -d -v ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/app/main cmd/cache/main.go \    
    && upx --lzma /opt/bin/app/main

FROM debian:stretch-slim AS runtime
COPY --from=builder /opt/bin/app /opt/bin/app

CMD [ "/opt/bin/app/main" ]
