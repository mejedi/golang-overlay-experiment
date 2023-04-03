build:
	go build -overlay overlay.json ./cmd/main

run:
	go run -overlay overlay.json ./cmd/main

test:
	go test -overlay overlay.json ./cmd/main
