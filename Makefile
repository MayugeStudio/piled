.PHONY: fmt test

piled.out: piled.go
	go build -o piled.out .

fmt:
	go fmt ./...

test: piled.out
	./test.py
	
