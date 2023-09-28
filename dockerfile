FROM golang:latest
WORKDIR /dev
COPY . .
ENV DOCKER_CONTAINER=true

RUN go mod download
RUN go build -o main .
EXPOSE 3000
CMD ["./main"]