FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY backend ./

WORKDIR /app/backend
RUN go mod download
COPY /backend .
RUN go build -o /app/backend/goalboard

FROM node:20.11-alpine AS frontend
WORKDIR /app
ADD /frontend ./
RUN npm install
RUN npm run build

FROM alpine
COPY --from=builder /app/backend/goalboard /app/backend
# COPY .env ./
COPY --from=frontend /app/dist /public

EXPOSE 8000
CMD [ "/app/backend" ]