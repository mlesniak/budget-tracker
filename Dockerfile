FROM golang:alpine
RUN apk add --no-cache alpine-sdk
RUN go get -v -u github.com/mlesniak/budget-tracker

FROM alpine:latest  
COPY --from=0 /go/bin/budget-tracker /usr/local/bin
ENTRYPOINT ["/usr/local/bin/budget-tracker"]