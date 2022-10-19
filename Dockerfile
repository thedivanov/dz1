FROM golang:1.19.2-alpine3.16 AS builder

WORKDIR /dz1
COPY . .

RUN go build -o /tmp/dz1 .


FROM alpine:3.16

EXPOSE 8000
WORKDIR /dz1

COPY --from=builder /tmp/dz1 .
ENTRYPOINT [ "./dz1" ]
