FROM golang:latest AS builder

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN mkdir -p /go/src/github.com/nsqio/nsq/apps/nsq_to_filebeat
ADD . /go/src/github.com/nsqio/nsq/apps/nsq_to_filebeat
WORKDIR /go/src/github.com/nsqio/nsq/apps/nsq_to_filebeat

RUN go mod download
RUN go mod verify

# Compile the binary and statically link
RUN CGO_ENABLED=0 go build -ldflags '-d -w -s'

RUN ls -lh

FROM elastic/filebeat:7.13.2
#FROM centos:7

ENV APP_DIR /data/
USER root
RUN mkdir $APP_DIR
COPY --from=builder /go/src/github.com/nsqio/nsq/apps/nsq_to_filebeat/nsq_to_filebeat $APP_DIR
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN chown -R filebeat:filebeat $APP_DIR

WORKDIR $APP_DIR
USER filebeat

