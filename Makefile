# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=echo-server

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
	docker build -t echo-server .

docker.run:
	docker run -p 8080:8080 echo-server
