FROM golang:1.17-alpine AS build
RUN apk add --no-cache git

WORKDIR /tmp/build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o ./out/bashdoard .

FROM alpine:3.9
COPY --from=build /tmp/build/out/bashdoard /app/bashdoard
EXPOSE 8080
CMD ["/app/bashdoard"]