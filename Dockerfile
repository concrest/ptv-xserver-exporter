#first stage - builder
FROM golang:1.12.1-stretch as builder

ARG BUILD_BUILDNUMBER=0.1-alpha
ARG BUILD_SOURCEVERSION=local-dev-build
ARG BUILD_DATE=unknown-date

COPY . /build
WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -ldflags "-X 'main.version=$BUILD_BUILDNUMBER' -X 'main.buildDate=$BUILD_DATE' -X 'main.commitHash=$BUILD_SOURCEVERSION'" -o ptv-xserver-exporter

#second stage
FROM alpine:latest
WORKDIR /app/
COPY --from=builder /build .
CMD ["/app/ptv-xserver-exporter"]