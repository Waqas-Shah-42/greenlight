FROM golang:1.20 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -o /greenlight ./cmd/api

FROM build-stage AS run-test-stage
RUN go test -v ./...

# Change the image to gcr.io/distroless/base-debian11:debug in order to ssh into the container
# docker run --rm -it --entrypoint sh <image_name>  
FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage greenlight greenlight
EXPOSE 4000
USER nonroot:nonroot
ENTRYPOINT ["./greenlight"]
