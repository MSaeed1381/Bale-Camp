FROM golang:alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /messenger-docker

EXPOSE 8000

CMD ["/messenger-docker"]
