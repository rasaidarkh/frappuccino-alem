.PHONY: start build-app fresh logs down

start:
	docker-compose up -d db
	docker-compose up -d app

build-app:
	docker-compose build --no-cache app

fresh: build-app
	docker-compose stop app
	docker-compose up -d app

logs:
	docker-compose logs -f app

down:
	docker-compose down
