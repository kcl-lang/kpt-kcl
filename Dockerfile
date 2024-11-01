FROM golang:1.23 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

ENV CGO_ENABLED=0
RUN GOOS=linux GOARCH=amd64 go build -o kpt-kcl-fn

FROM kcllang/kcl

WORKDIR /app

COPY --from=builder /app/kpt-kcl-fn .

CMD ["/app/kpt-kcl-fn"]
