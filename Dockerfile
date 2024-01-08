FROM golang:1.21-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o gpgencryptor .

FROM alpine:3.19 as app

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
