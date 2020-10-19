all: build

build:
	docker build . -t hekonsek/dkr-docker-last-id
	docker push hekonsek/dkr-docker-last-id