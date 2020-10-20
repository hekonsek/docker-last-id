all: build

build:
	go build docker-last-id.go
	docker build . -t hekonsek/dkr-docker-last-id
	docker push hekonsek/dkr-docker-last-id