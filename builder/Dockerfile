FROM golang:1.19 as build-env
ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN go mod tidy -v
RUN go mod download && go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo -o /dist/todo-app-fiber main.go
RUN chmod +x /dist/todo-app-fiber

FROM alpine:latest
RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta
COPY --from=build-env /dist/todo-app-fiber app
ENTRYPOINT ["./app"]