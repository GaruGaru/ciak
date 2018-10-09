VERSION=$(shell git rev-parse --short HEAD)
DIST_DIR=./dist

# GO

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build:
	go build -o ${DIST_DIR}/ciak ciak.go

.PHONY: run
run:
	go run ciak.go

.PHONY: test
test:
	go test -v ./...

.PHONY: deps
deps:
	dep ensure



# DOCKER

IMAGE=garugaru/ciak:${VERSION}

BASE_COMPOSE=docker/docker-compose.yml

.PHONY: docker-build
docker-build:
	docker build -f docker/Dockerfile -t ${IMAGE} .

.PHONY: docker-push-image
docker-push-image: docker-build
	docker push ${IMAGE}


.PHONY: docker-up
docker-up:
	docker-compose -f ${BASE_COMPOSE} up

.PHONY: docker-upd
docker-upd:
	docker-compose -f ${BASE_COMPOSE} up -d


.PHONY: docker-down
docker-down:
	docker-compose -f ${BASE_COMPOSE} down