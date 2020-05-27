# Build stage
FROM golang:1.14.3-alpine3.11 AS builder

RUN apk add librdkafka-dev git build-base && \
      rm -rf /var/cache/apk/*

# Enable support of go modules by default
ENV GO111MODULE on
ENV BASE_DIR /go/src/pipeline-endpoint

# Warming modules cache with project dependencies
WORKDIR ${BASE_DIR}
COPY go.mod go.sum ./
RUN go mod download

# Copy project source code to WORKDIR
COPY . .

# Run tests and build on success
RUN go test ./... \
 && go build -o /go/bin/pipeline-endpoint


# Final container stage
FROM golang:1.14.3-alpine3.11

# install rdkafka
RUN apk update && \
      apk add librdkafka-dev \
      ca-certificates && \
      rm -rf /var/cache/apk/*

COPY --from=builder /go/bin/pipeline-endpoint /bin/pipeline-endpoint
