FROM golang:1.21.7-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN apk add --no-cache tzdata

RUN go mod download

RUN go build -o /goapp

FROM alpine:3.19 AS build-release-stage

WORKDIR /

COPY --from=build-stage /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build-stage /goapp /goapp
COPY --from=build-stage /app/templates ./templates

EXPOSE 8000

ENTRYPOINT ["./goapp"]