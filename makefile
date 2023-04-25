VERSION:=$(shell cat VERSION)

default: run

run:
	kpt fn source ./testdata/resources.yaml --fn-config ./testdata/fn-config.yaml | go run main.go

test:
	go test ./...

fmt:
	go fmt ./...		

image:

image:
	docker build . -t docker.io/peefyxpf/kpt-kcl:$(VERSION)
	docker push docker.io/peefyxpf/kpt-kcl:$(VERSION)
