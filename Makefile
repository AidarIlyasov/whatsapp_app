build:
	@go build -o ./bin/whatsapp cmd/main.go
run: build
	./bin/whatsapp