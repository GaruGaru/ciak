ARG GO_VERSION=1.11
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

FROM jrottenberg/ffmpeg:4.0-alpine AS final

WORKDIR /

RUN mkdir /data && mkdir /transfer

VOLUME /transfer

VOLUME /data

COPY static/ /static

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app /app

EXPOSE ${PORT}

ENTRYPOINT ["/app"]