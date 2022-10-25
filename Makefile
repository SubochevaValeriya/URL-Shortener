build:
	docker-compose build shortener-app

run:
	docker-compose up shortener-app

test:
	go test -v ./...