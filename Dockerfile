FROM golang:1.22

RUN mkdir /app
WORKDIR /app
COPY . .