FROM golang:1.17.2 as builder
RUN mkdir -p /go/src/
WORKDIR /go/src/
COPY . /go/src/
RUN go get -d
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o location .

FROM alpine:latest
WORKDIR /
COPY --from=builder /go/src/location .
ENV DEFAULT_PORT=3000
ENV DNS=root:secret@/location
CMD ["/location","listen"]