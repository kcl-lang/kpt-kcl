FROM golang:1.19 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o kpt-kcl-fn


FROM kusionstack/kclvm

WORKDIR /app
USER root
COPY --from=builder /app/kpt-kcl-fn .
RUN mkdir -p /root/go/bin
RUN echo "latest" > /root/go/bin/kclvm.version
RUN rm -rf /kclvm

CMD ["/app/kpt-kcl-fn"]
