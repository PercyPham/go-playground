FROM golang:1.12.7-alpine
WORKDIR /app
COPY src ./src
CMD ["go", "run", "src/main.go"]