FROM golang:1.16-alpine

RUN apk update

RUN apk upgrade

RUN apk add make

RUN add gcc musl-dev 

COPY . ./project/order

