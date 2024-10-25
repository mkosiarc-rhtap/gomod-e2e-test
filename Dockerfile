FROM golang:1.20-alpine

# Test disabled network access
RUN if curl -IsS www.google.com; then echo "Has network access!"; exit 1; fi

WORKDIR /app


COPY . .
RUN go mod download
RUN go build -o myapp .

CMD ["./myapp"]
