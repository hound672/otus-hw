FROM golang:1.21 as builder

WORKDIR /app

# Кэшируем слои с модулями
COPY go.mod go.sum ./
RUN go mod download

COPY ../build ${CODE_DIR}

COPY build ./build
COPY pkg ./pkg
COPY cmd ./cmd
COPY internal ./internal

ARG VERSION

RUN CGO_ENABLED=0 -ldflags="-X github.com/hound672/otus-hw/hw12_13_14_15_calendar/build.Version=$(VERSION)" -v -o calendar ./cmd

FROM alpine:3.18.4

WORKDIR /app

COPY config.yml ./
COPY --from=builder /app/calendar ./

CMD ["/app/calendar"]
