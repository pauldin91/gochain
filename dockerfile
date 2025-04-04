FROM golang:1.24 AS build

WORKDIR /application

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./src/main.go 

FROM alpine:3.21

WORKDIR /application

COPY --from=build /application/* .

COPY certificates ./certificates
COPY docs ./docs

CMD ["./main"]


