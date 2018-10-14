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

IMAGE=garugaru/ciak
ARMHF_IMAGE=garugaru/rpi-ciak

BASE_COMPOSE=docker/docker-compose.yml

.PHONY: docker-build
docker-build:
	docker build -f docker/Dockerfile -t ${IMAGE}:latest -t ${IMAGE}:${VERSION} .

.PHONY: docker-push-image
docker-push-image: docker-build
	docker push ${IMAGE}:${VERSION}

.PHONY: docker-build-arm
docker-build-arm:
	docker build -f docker/Dockerfile.armhf -t ${ARMHF_IMAGE}:latest -t ${ARMHF_IMAGE}:${VERSION} .

.PHONY: docker-push-image-arm
docker-push-image-arm: docker-build-arm
	docker push ${ARMHF_IMAGE}:${VERSION}

.PHONY: docker-up
docker-up:
	docker-compose -f ${BASE_COMPOSE} up

.PHONY: docker-upd
docker-upd:
	docker-compose -f ${BASE_COMPOSE} up -d


.PHONY: docker-down
docker-down:
	docker-compose -f ${BASE_COMPOSE} down