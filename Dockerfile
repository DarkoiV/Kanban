# BUILD BACKEND
FROM golang:1.16-alpine AS go-builder

WORKDIR /go/src/backend
COPY backend .

RUN apk add build-base
RUN go build -o server 

# BUILD FRONTEND
FROM node:16-alpine AS npm-builder

WORKDIR /frontend
COPY frontend .

RUN npm install 
RUN npm run build

# FINAL IMAGE
FROM alpine

WORKDIR /app

COPY .env .
COPY --from=go-builder /go/src/backend/server .
COPY --from=npm-builder /frontend/dist ./dist

cmd ["./server"]
