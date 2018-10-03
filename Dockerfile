FROM golang:alpine as builder
RUN mkdir /build 
RUN apk add --no-cache git  \
    && go get github.com/PuerkitoBio/goquery \
    && go get github.com/gorilla/mux \
    && go get github.com/gorilla/handlers \
    && apk del git 
ADD . /build/
WORKDIR /build 
RUN go build -o main .
FROM alpine
RUN apk --no-cache add ca-certificates
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
CMD ["./main"]