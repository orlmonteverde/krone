default:
	go build .

test:
	go test -v ./...

cover:
	go test -cover -v ./...
