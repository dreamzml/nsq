FROM 10.32.5.83:32500/golang:1.16.4 as builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN mkdir -p /go/src/github.com/nsqio/nsq/apps/producer_test
WORKDIR /go/src/github.com/nsqio/nsq/apps/producer_test
ADD . ./

RUN go mod download
RUN go mod verify

# Compile the binary and statically link
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s'

FROM 10.32.5.83:32500/alpine:3.12.0 as producer-test

ENV APP_DIR /data/
WORKDIR $APP_DIR

COPY --from=builder /go/src/github.com/nsqio/nsq/apps/producer_test/producer_test $APP_DIR
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ENTRYPOINT ["./producer_test"]