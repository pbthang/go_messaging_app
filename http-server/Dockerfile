FROM golang:1.18

RUN apt-get update && apt-get install -y netcat

WORKDIR /app
COPY . .
RUN go build -o main

EXPOSE 8080
CMD ["./main"]
