compose-rebuild:
	docker compose up -d --no-deps --build service

compose-rebuild-alt:
	docker compose up -d --no-deps --build service

service-build:
	go build -o ./build/service ./cmd/main.go

