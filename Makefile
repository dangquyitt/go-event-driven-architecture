.PHONY: server
server:
	go run cmd/api/main.go

.PHONY: docker-compose-up
docker-compose-up:
	docker-compose -f docker-compose.yml up -d

.PHONY: docker-compose-down
docker-compose-down:
	docker-compose -f docker-compose.yml down

	