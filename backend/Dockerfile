FROM golang:1.22-bookworm AS builder

WORKDIR /src

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch

COPY --from=builder /app /app
COPY --from=builder /src/db.sqlite /

EXPOSE 8080

ENTRYPOINT ["/app"]
