VERSION:=$(shell cat VERSION)

default: run

run:
	kpt fn source ./testdata/resources.yaml --fn-config ./testdata/fn-config.yaml | go run main.go

test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/kcllang/kpt-kcl:v$(VERSION)
	docker push docker.io/kcllang/kpt-kcl:v$(VERSION)

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)
	gh release create v$(VERSION) --draft --generate-notes --title "$(VERSION) Release"
