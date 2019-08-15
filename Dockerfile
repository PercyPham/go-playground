FROM golang:1.12.8-alpine
WORKDIR /app
COPY src ./src
CMD ["go", "run", "src/main.go"]