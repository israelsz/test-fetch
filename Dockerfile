FROM golang:1.19-alpine

RUN apk --no-cache add ca-certificates

# Set destination folder
WORKDIR /app/
# Copy and Download Go modules of the project
COPY go.mod .
RUN go mod download

# Copy the source code inside the container.
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /pingeso-fusupo-backend-v2

# Port to be exposed
EXPOSE 8000

# Run the app
CMD ["/pingeso-fusupo-backend-v2"]