FROM golang:alpine
RUN apk add --no-cache alpine-sdk

# Prevent caching, see https://stackoverflow.com/questions/31782220/how-can-i-prevent-a-dockerfile-instruction-from-being-cached
ARG CACHE_DATE=not_a_date
RUN go get -v -u github.com/mlesniak/budget-tracker

FROM alpine:latest  
COPY --from=0 /go/bin/budget-tracker /usr/local/bin
ENTRYPOINT ["/usr/local/bin/budget-tracker"]