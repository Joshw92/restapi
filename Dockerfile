FROM golang:1.11 as builder
WORKDIR /go/src/github.com/Joshw92/
RUN go get -d -v golang.org/x/net/html  
RUN go get -d github.com/gorilla/mux
RUN go get -u github.com/lib/pq
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /go/src/github.com/Joshw92/app .
CMD ["./app"]  
