FROM golang:alpine
WORKDIR /go/src/app
COPY . .
RUN go install -v ./...
RUN ls /go/bin

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/bin/app .
CMD ["/root/app"]