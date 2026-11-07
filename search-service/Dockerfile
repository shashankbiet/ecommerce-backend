FROM golang:1.26-alpine AS build-stage
RUN apk add --no-progress --no-cache gcc musl-dev

WORKDIR /build

COPY search-service/go.mod search-service/go.sum ./
RUN go mod download

COPY search-service/. .

RUN --mount=type=cache,id=search-go-build,target=/root/.cache/go-build \
    GOMAXPROCS=1 go build -p=1 -tags musl -ldflags '-extldflags "-static"' -o /build/bin/app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12:nonroot AS build-release-stage

WORKDIR /src

COPY --from=build-stage /build/bin/app /src/app
COPY --from=build-stage /build/conf /src/conf

USER nonroot:nonroot

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 CMD ["/src/app", "healthcheck"]

ENTRYPOINT ["/src/app"]