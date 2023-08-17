FROM golang:1.20-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/go-sample-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


RUN go build -ldflags="-s -w" -o ./out/go-sample-app ./cmd/main.go

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-sample-app/out/go-sample-app /app/go-sample-app
COPY ./block.txt .
EXPOSE 3000

CMD ["/app/go-sample-app"]