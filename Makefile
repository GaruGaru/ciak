BIN_NAME=ciak
BIN_OUTPUT=dist/${BIN_NAME}

fmt:
	go fmt ./...

deps:
	go mod vendor
	go mod verify

build: fmt deps
	go build -o ${BIN_OUTPUT}


DOCKER_IMAGE=garugaru/ciak
COMPOSE=docker/docker-compose.yml
VERSION=$(shell git rev-parse --short HEAD)
DOCKERFILE_ARMHF=Dockerfile.armhf

docker-up:
	docker-compose -f ${COMPOSE} up

docker-upd:
	docker-compose -f ${COMPOSE} up -d

docker-down:
	docker-compose -f ${COMPOSE} down

docker-build:
	docker build -t ${DOCKER_IMAGE}:${VERSION} .

docker-push: docker-build
	docker push ${DOCKER_IMAGE}:${VERSION}

docker-build-arm:
	docker build -t ${DOCKER_IMAGE}-armhf:${VERSION} -f ${DOCKERFILE_ARMHF} .

docker-push-arm: docker-build-arm
	docker push ${DOCKER_IMAGE}-armhf:${VERSION}
