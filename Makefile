GO_OUTPUT ?= go-simple-api
GO_EXPOSE_PORT ?= 3000
DOCKER_IMAGE_NAME ?= aassuhendar/go-simple-api
DOCKER_IMAGE_VERSION ?= latest

go-build:
	dep ensure
	CGO_ENABLED=0 GOOS=linux go build -a -o ./builds/$(GO_OUTPUT) main.go

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) .

docker-run:
	docker run -d -p $(GO_EXPOSE_PORT):$(GO_EXPOSE_PORT) --name $(GO_OUTPUT) --rm $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

docker-stop:
	docker stop $(GO_OUTPUT)

docker-logs:
	docker logs $(GO_OUTPUT)

clean-build:
	rm -f ./builds/$(GO_OUTPUT)

clean-docker:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION)

clean-all:
	docker rmi -f $(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_VERSION) && rm -f ./builds/$(GO_OUTPUT)
