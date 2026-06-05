FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/portfolio-ssh ./cmd/server

FROM alpine:3.22

RUN adduser -D -h /app appuser

WORKDIR /app
RUN mkdir -p /app/.ssh && chown -R appuser:appuser /app

COPY --from=build /out/portfolio-ssh /usr/local/bin/portfolio-ssh

USER appuser

EXPOSE 2222

ENV HOST=0.0.0.0
ENV PORT=2222
ENV HOST_KEY_PATH=/app/.ssh/ssh_host_ed25519_key

ENTRYPOINT ["portfolio-ssh"]
