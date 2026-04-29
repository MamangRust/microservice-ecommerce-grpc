FROM golang:1.25.3-bookworm AS builder

WORKDIR /app

# Copy dependency manifests first (cache friendly)
COPY shared/go.mod shared/go.sum ./shared/
COPY pkg/go.mod pkg/go.sum ./pkg/
COPY service/merchant_business/go.mod service/merchant_business/go.sum ./service/merchant_business/

WORKDIR /app/service/merchant_business

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

WORKDIR /app

# Copy source
COPY shared/ ./shared/
COPY pkg/ ./pkg/
COPY service/merchant_business/ ./service/merchant_business/

WORKDIR /app/service/merchant_business

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/merchant_business_bin ./cmd/main.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1000 appuser && \
    adduser -D -s /bin/sh -u 1000 -G appuser appuser

WORKDIR /app

RUN mkdir -p /var/log/app && chown -R 1000:1000 /var/log/app

COPY --from=builder /app/merchant_business_bin ./merchant_business

RUN chown -R 1000:1000 ./merchant_business

USER appuser

EXPOSE 50066
CMD ["./merchant_business"]
