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
DOCKER_IMAGE_ARM=${DOCKER_IMAGE}:armhf
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
	docker build -t ${DOCKER_IMAGE}:latest -t ${DOCKER_IMAGE}:${VERSION} .

docker-push: docker-build
	docker push ${DOCKER_IMAGE}:amd64-${VERSION}
	docker push ${DOCKER_IMAGE}:amd64

docker-build-arm:
	docker build -t ${DOCKER_IMAGE}:arm -t ${DOCKER_IMAGE}:arm-${VERSION} -f ${DOCKERFILE_ARMHF} .

docker-push-all: docker-build-arm docker-build
	docker push ${DOCKER_IMAGE}:amd64-${VERSION}
	docker push ${DOCKER_IMAGE}:amd64
	docker push ${DOCKER_IMAGE}:arm
	docker push ${DOCKER_IMAGE}:arm-${VERSION}

	docker manifest create garugaru/ciak ${DOCKER_IMAGE}:amd64 ${DOCKER_IMAGE}:amd64-${VERSION} ${DOCKER_IMAGE}:arm ${DOCKER_IMAGE}:arm-${VERSION}

	docker manifest annotate --arch arm garugaru/ciak ${DOCKER_IMAGE}:arm
	docker manifest annotate --arch arm garugaru/ciak ${DOCKER_IMAGE}:arm-${VERSION}
	docker manifest push garugaru/ciak

docker-create-manifest:
	docker manifest create ${DOCKER_IMAGE}:latest ${DOCKER_IMAGE}:latest ${DOCKER_IMAGE}:arm-latest



