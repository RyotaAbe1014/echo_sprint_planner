FROM golang:latest

WORKDIR /app
COPY ./app /app
COPY ../go.mod /app
COPY ../go.sum /app

RUN go mod tidy \
    && go mod download

ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64
EXPOSE 8080


# CMD ["go", "run", "main.go"]