default: run

run:
	kpt fn source ./testdata/resources.yaml --fn-config ./testdata/fn-config.yaml | go run main.go

test:
	go test ./...

fmt:
	go fmt ./...		
