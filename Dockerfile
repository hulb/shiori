# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache build-base git
WORKDIR /src
COPY . .
RUN go build -ldflags "-s -w"

# server image
FROM alpine:latest
COPY --from=builder /src/shiori /usr/local/bin/
ENV SHIORI_DIR /srv/shiori/
EXPOSE 8080
CMD ["/usr/local/bin/shiori", "serve"]
