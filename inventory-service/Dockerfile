FROM golang:1.25-alpine AS build-stage
RUN apk add --no-progress --no-cache gcc musl-dev

WORKDIR /build

COPY inventory-service/go.mod inventory-service/go.sum ./
RUN go mod download

COPY inventory-service/. .

RUN --mount=type=cache,id=inventory-go-build,target=/root/.cache/go-build \
    go build -tags musl -ldflags '-extldflags "-static"' -o /build/bin/app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /src

COPY --from=build-stage /build/bin/app /src/app
COPY --from=build-stage /build/conf /src/conf

ENTRYPOINT ["/src/app"]