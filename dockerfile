FROM golang:alpine AS builder
WORKDIR /go/ai4u
COPY . /go/ai4u
ENV DOCKER_CONTAINER=true

RUN go mod download
RUN GOOS=linux  CGO_ENABLED=0 go build -ldflags="-w -s" main.go

FROM alpine AS runner
WORKDIR /go/ai4u
COPY --from=builder /go/ai4u/main .
COPY --from=builder /go/ai4u/go.mod .
COPY --from=builder /go/ai4u/go.sum .
COPY --from=builder /go/ai4u/appsettings.json .
COPY --from=builder /go/ai4u/cert.pem .
COPY --from=builder /go/ai4u/key.pem .
COPY --from=builder /go/ai4u/templatesite ./templatesite
EXPOSE 3000
ENTRYPOINT ["./main"]