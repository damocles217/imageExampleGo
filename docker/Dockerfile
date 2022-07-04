FROM golang:latest

EXPOSE 7700

WORKDIR /usr/app/

COPY ../ .

RUN CGO_ENABLED=0 GOOS=linux go build -o main

CMD ["make"]

FROM alpine:latest

ENTRYPOINT [ "./server" ]
