build:
	docker-compose build shortener-app

run:
	docker-compose up shortener-app

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go