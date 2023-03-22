FROM golang:1.18 as builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o kpt-kcl-fn


FROM kusionstack/kclvm

WORKDIR /app

COPY --from=builder /app/kpt-kcl-fn .
RUN mkdir -p /go/bin

CMD ["/app/kpt-kcl-fn"]
