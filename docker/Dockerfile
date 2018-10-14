FROM golang:1.10 AS build

RUN mkdir -p $GOPATH/src/github.com/GaruGaru/ciak
ADD . $GOPATH/src/github.com/GaruGaru/ciak
WORKDIR $GOPATH/src/github.com/GaruGaru/ciak
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN CGO_ENABLED=0 go build -o ciak -ldflags="-s -w" ciak.go
RUN cp ciak /

FROM jrottenberg/ffmpeg:4.0-alpine

VOLUME /data

COPY static/ /home/static/
COPY --from=build /ciak /home/ciak
RUN ln -s /home/ciak /usr/bin/ciak
WORKDIR /home
ENTRYPOINT ["ciak"]