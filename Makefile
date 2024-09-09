.PHONY: fmt test

piled.out: main.go
	go build -o piled.out .

fmt:
	go fmt ./...

cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	explorer.exe cover.html

test: piled.out
	./test.py
	
