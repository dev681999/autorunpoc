FROM golang:1.14-alpine as backend_builder

RUN apk add upx

WORKDIR /app

COPY . .

RUN go build -o server -ldflags="-s -w"
RUN upx -q -5 server

FROM alpine:3.9

WORKDIR /app
COPY --from=backend_builder /app/server /

EXPOSE 8090

CMD [ "/server" ]
