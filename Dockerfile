ARG GO_VERSION=1.15
ARG APP_NAME="ciak"
ARG PORT=8082

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -ldflags="-s -w" \
    -installsuffix 'static' \
    -o /app .

FROM linuxserver/ffmpeg

WORKDIR /

RUN mkdir /data && mkdir /transfer

VOLUME /transfer

VOLUME /data

VOLUME /db

COPY ui/ /ui

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app /app

EXPOSE ${PORT}

ENTRYPOINT ["/app"]