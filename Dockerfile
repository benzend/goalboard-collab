FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
ADD /backend ./
COPY backend/go.mod ./backend
COPY backend/go.sum ./backend

WORKDIR /app/backend
RUN go mod download
COPY /backend .
RUN go build -o /app/backend

FROM node:20.11-alpine AS frontend
WORKDIR /app
ADD /frontend ./
RUN npm install
RUN npm run build

FROM alpine
COPY --from=builder /app /app
COPY .env ./
COPY --from=frontend /app/dist /public

EXPOSE 8000
CMD [ "/app" ]