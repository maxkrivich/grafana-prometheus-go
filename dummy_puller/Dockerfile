FROM golang:1.14-alpine as builder 
COPY . /dummy_puller
WORKDIR /dummy_puller
RUN CGO_ENABLED=0 \ 
    GOOS=linux \
    GO111MODULE=on \
    go build -mod=mod -o puller main.go
RUN chmod +x puller

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /dummy_puller .
CMD [ "./puller" ]