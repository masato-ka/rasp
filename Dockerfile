FROM golang:latest

MAINTAINER masato-ka

WORKDIR /go
RUN mkdir .go
ENV GOPATH /go/.go
ADD . /go

RUN go build

