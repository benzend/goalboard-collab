FROM golang:alpine AS base
RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . ./backend

FROM base AS development

CMD [ "go", "run", "./backend" ]

FROM base AS dev-envs

EXPOSE 8000

FROM base AS builder

RUN go build -o main

FROM base

ENV env=PRODUCTION

COPY --from=builder /app/backend/main /main

CMD [ "/main" ]