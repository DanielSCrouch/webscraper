run:
	@go build -o web-crawler ./cmd/main.go
	./web-crawler scrape "https://en.wikipedia.org/wiki/Ken_Thompson"

test:
	@go test ./...