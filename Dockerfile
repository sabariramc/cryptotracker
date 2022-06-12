FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
RUN apk add build-base

WORKDIR /myapp
COPY go.mod .
RUN go mod download

COPY . .
WORKDIR /myapp/src/cmd/server
RUN go build -o /app 


##
## Deploy
##
FROM alpine:latest

COPY --from=builder /app /app
RUN apk --no-cache add tzdata bash
EXPOSE 3000

ENTRYPOINT ["/app"]