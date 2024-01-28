FROM golang:1.21-alpine AS build-stage
RUN apk add --no-progress --no-cache gcc musl-dev

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -tags musl -ldflags '-extldflags "-static"' -o /build/bin/app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /src

COPY --from=build-stage /build/bin/app /src/app
COPY --from=build-stage /build/conf /src/conf

ENTRYPOINT ["/src/app"]