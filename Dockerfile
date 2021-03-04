FROM golang:1.16.0-alpine3.13 as builder  

RUN apk update && \
    apk upgrade && \
    apk --update add git 

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go get -v github.com/labstack/echo/v4 github.com/labstack/echo/v4/middleware
RUN go build -o app .

EXPOSE 80

CMD ["/app/app"]