FROM golang:1.19 AS build-stage

WORKDIR /inventory-service

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /inventory-service/bin/app

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /inventory-service

COPY --from=build-stage /inventory-service/bin/app /inventory-service/app
COPY --from=build-stage /inventory-service/conf /inventory-service/conf

ENTRYPOINT ["/inventory-service/app"]