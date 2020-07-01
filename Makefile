run:
	docker-compose up

test:
	docker build . --target test

build:
	docker build . --target bin