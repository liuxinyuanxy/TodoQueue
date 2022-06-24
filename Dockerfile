FROM golang:1.18 AS builder

ENV GOPROXY=https://goproxy.cn,direct \
    GO111MODULE=on \
    WORKDIR=/todoq/src \
    CGO_ENABLED=0

RUN mkdir -p $WORKDIR

COPY . $WORKDIR

RUN cd $WORKDIR && go mod download

RUN cd $WORKDIR && go install github.com/swaggo/swag/cmd/swag@latest && swag init

RUN cd $WORKDIR && go build -o /TodoQueue

FROM alpine:3.15.2

COPY ./conf.yaml /

COPY --from=builder /TodoQueue /TodoQueue

EXPOSE 1323

CMD ["/TodoQueue"]