
# Document that the service listens on port 8080.

FROM golang:1.9 as builder  
RUN mkdir -p /go/src/go_api   
WORKDIR /go/src/go_api    
COPY . .  
RUN go get -d  
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .    



FROM alpine:latest  
WORKDIR /
COPY --from=builder /go/src/go_api/app .  
EXPOSE 3000
CMD ["/app"]