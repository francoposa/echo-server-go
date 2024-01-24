# git version --dirty tag ensures we don't override an image tag built from a clean state
# with an image tag built from a "dirty" state with uncommitted changes.
# we should be able to use git to see the repo in the exact state the container was built from.
GIT_VERSION ?= $(shell git describe --abbrev=8 --tags --always --dirty)
IMAGE_PREFIX ?= ghcr.io/francoposa/echo-server-go/echo-server
SERVICE_NAME ?= echo-server

.PHONY: clean
clean:
	rm -rf dist/

.PHONY: local.build
local.build: clean
	go build -o dist/echo-server ./src/cmd/server

.PHONY: local.test
local.test:
	go test -v ./...

.PHONY: local.run
local.run:
	go run ./src/cmd/server/main.go

.PHONY: docker.build
docker.build:
	# image gets tagged as latest by default
	docker build -t $(IMAGE_PREFIX)/$(SERVICE_NAME) -f ./server/Dockerfile .
	# tag with git version as well
	docker tag $(IMAGE_PREFIX)/$(SERVICE_NAME) $(IMAGE_PREFIX)/$(SERVICE_NAME):$(GIT_VERSION)

.PHONY: docker.run
docker.run:  # defaults to latest tag
	docker run -p 8080:8080 $(IMAGE_PREFIX)/$(SERVICE_NAME)

.PHONY: docker.push
docker.push: docker.build
	docker push $(IMAGE_PREFIX)/$(SERVICE_NAME):$(GIT_VERSION)
	docker push $(IMAGE_PREFIX)/$(SERVICE_NAME):latest