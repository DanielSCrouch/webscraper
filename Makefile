run:
	@go build -o web-crawler .
	./web-crawler

test:
	@go test ./...