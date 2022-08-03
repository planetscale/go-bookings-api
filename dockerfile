FROM golang:1.18-bullseye

RUN go install golang.org/dl/go1.18@latest \
  && go1.18 download

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./server

CMD ["/app/server"]