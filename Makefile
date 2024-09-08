piled.out: main.go asm/generator.go asm/lexer.go
	go build -o piled.out .

fmt:
	go fmt ./...

cover:
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html
	explorer.exe cover.html

ex: piled.out
	./piled.out examples/basic-inst.piled
	
