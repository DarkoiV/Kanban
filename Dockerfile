FROM debian:latest

WORKDIR /app
COPY ./build/latest/ /app

CMD ["./server"]
