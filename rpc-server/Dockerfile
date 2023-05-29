FROM golang:1.18

# Install netcat
RUN apt-get update && apt-get install -y netcat

WORKDIR /app
COPY . .
RUN sh ./build.sh

EXPOSE 8888
CMD ["./output/bootstrap.sh"]
