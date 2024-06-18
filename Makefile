generate:
	go generate ./...

docker-build:
	docker build -t go-blog-service .

docker-run:
	docker run --rm -p 8080:8080 go-blog-service

run:
	go run cmd/server/main.go

test:
	go test ./...
