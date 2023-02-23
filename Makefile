winbuild:
	GOOS=windows GOARCH=amd64 go build

build:
	go build

test:
	go test ./...

testshort:
	go test ./... -short

vet: 
	go vet ./...
	
buildandtest: build test

run:
	go run .

upgrade:
	go get -u