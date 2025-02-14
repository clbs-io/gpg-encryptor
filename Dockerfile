# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
ARG TARGETOS TARGETARCH
ARG VERSION=v0.0.0

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o gpgencryptor .

FROM alpine:latest AS gpg-encryptor

ARG USERNAME="gpgencryptor"
ARG UID=1001
ARG GID=1001

RUN addgroup \
    -g ${GID} \
    ${USERNAME} \
    && \
    adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "${USERNAME}" \
    --uid "$UID" \
    "${USERNAME}"

WORKDIR /app

COPY --from=builder /build/gpgencryptor .

USER ${USERNAME}

CMD [ "/app/gpgencryptor" ]
