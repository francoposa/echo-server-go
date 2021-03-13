# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=echo-server

# Docker parameters
#DOCKER_REGISTRY=docker.pkg.github.com
DOCKER_REGISTRY=ghcr.io
DOCKER_ORG=francoposa
DOCKER_REPO=echo-server-go
DOCKER_IMAGE=echo-server
DOCKER_DEFAULT_TAG=latest
DOCKER_FULL_IMAGE_NAME=$(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_REPO)/$(DOCKER_IMAGE):$(DOCKER_DEFAULT_TAG)


local.build:
	$(GOBUILD) -o $(BINARY_NAME)
local.test:
	$(GOTEST) ./...
local.clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
local.run:
	$(GORUN) ./ server --config config.local.yaml

docker.build:
	docker build -t $(DOCKER_FULL_IMAGE_NAME) .

docker.push:
	docker push $(DOCKER_FULL_IMAGE_NAME)

docker.run:
	docker run -p 8080:8080 $(DOCKER_FULL_IMAGE_NAME)

