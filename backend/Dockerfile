FROM golang:alpine AS base
RUN apk update && apk add --no-cache git

RUN apk add postgresql-client

WORKDIR /code

COPY go.mod go.sum ./
RUN go mod download

ADD . .

FROM base AS development

CMD [ "go", "run", "." ]

FROM base AS dev-envs

EXPOSE 8000

FROM base AS builder

RUN go build -o main

FROM base

ENV env=PRODUCTION

COPY --from=builder /code/main /main

CMD [ "/main" ]