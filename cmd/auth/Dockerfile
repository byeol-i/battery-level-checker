FROM golang:1.18 AS builder

COPY . "/go/src/github.com/byeol-i/battery-level-checker/"
WORKDIR "/go/src/github.com/byeol-i/battery-level-checker/"

RUN mkdir -p /opt/bin/app/

RUN go mod vendor \
    && go get -u -d -v ./...
    
RUN go build -mod=mod -o /opt/bin/app/ ./cmd/auth 

FROM debian:stretch-slim AS runtime
RUN apt update && apt install -y ca-certificates 

COPY --from=builder /opt/bin/app /opt/bin/app

CMD [ "/opt/bin/app/auth" ]
