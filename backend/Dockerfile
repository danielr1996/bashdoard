FROM golang:1.18-alpine AS build
RUN apk add --no-cache git

WORKDIR /tmp/build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/bashdoard .

FROM alpine:3.9
COPY --from=build /tmp/build/out/bashdoard /app/bashdoard
EXPOSE 3040
CMD ["/app/bashdoard"]