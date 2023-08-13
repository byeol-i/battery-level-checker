FROM golang:1.20-alpine AS builder

# RUN apt update && apt install -y ca-certificates upx
RUN apk add upx ca-certificates

WORKDIR "/go/src/github.com/byeol-i/battery-level-checker/"

COPY go.mod .
COPY go.sum .

RUN mkdir -p /opt/bin/app/

# COPY vendor/ /go/src/

COPY . .

RUN go mod tidy \
    && go get -u -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w' -o /opt/bin/app/main cmd/apid/main.go \    
    && upx --lzma /opt/bin/app/main

FROM debian:stretch-slim AS runtime
COPY --from=builder /opt/bin/app /opt/bin/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD [ "/opt/bin/app/main" ]
