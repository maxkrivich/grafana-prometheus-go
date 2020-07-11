FROM golang:1.14-alpine as builder
COPY . /dummy_pusher
WORKDIR /dummy_pusher
RUN CGO_ENABLED=0 \ 
    GOOS=linux \
    GO111MODULE=on \
    go build -mod=mod -o pusher main.go
RUN chmod +x pusher

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /dummy_pusher .
CMD [ "./pusher" ]