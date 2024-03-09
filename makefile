VERSION:=$(shell cat VERSION)

default: run

run:
	kpt fn source ./testdata/resources.yaml --fn-config ./testdata/fn-config.yaml | go run main.go

test:
	go test ./...

fmt:
	go fmt ./...		

image:
	docker build . -t docker.io/kcllang/kpt-kcl:$(VERSION)
	docker push docker.io/kcllang/kpt-kcl:$(VERSION)

release:
	git tag $(VERSION)
	git push origin $(VERSION)
	gh release create $(VERSION) --draft --generate-notes --title "$(VERSION) Release"
