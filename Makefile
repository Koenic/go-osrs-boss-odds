build:
	go build -o bin/rs-odds cmd/rs-odds/main.go

test:
	go test ./...

graph:
	python3 python/make-graph.py

run:
	go run cmd/rs-odds/main.go

profile:
	go run cmd/rs-odds/main.go test