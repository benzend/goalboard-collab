FROM golang:alpine AS base
RUN apk update && apk add --no-cache git bash

RUN apk add postgresql-client

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

ADD . .

FROM base AS test

ENV env=TEST

FROM base AS development

ENV env=DEVELOPMENT

CMD [ "go", "run", "." ]

FROM base AS dev-envs

EXPOSE 8000

FROM base AS builder

RUN go build -o main

FROM base

ENV env=PRODUCTION

COPY --from=builder /code/main /main

CMD [ "/main" ]