FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
COPY . /xss
WORKDIR /xss
RUN go build -o xss cmd/lab-xss/main.go

FROM alpine:latest
COPY --from=builder /xss/xss .
COPY --from=builder /xss/static static
ENTRYPOINT [ "./xss" ]
CMD ["help"]
