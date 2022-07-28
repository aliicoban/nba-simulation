FROM golang:1.18

RUN go version
ENV GOPATH=/

COPY ./ ./



# build go app
RUN go mod download

RUN go build -o nba-simulation ./cmd/main.go

# Expose application port
EXPOSE 4444

CMD ["./nba-simulation]