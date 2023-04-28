FROM golang

# Set destination folder
WORKDIR /go/src/fusupo-backend/
# Copy and Download Go modules of the project
COPY go.mod .
RUN go mod download

# Copy the source code inside the container.
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /pingeso-fusupo-backend-v2

######## New stage #######
FROM alpine:3.14  

RUN apk update && apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the folder with the binary file from the previous stage
COPY --from=0 /go/src/fusupo-backend/ .

# Port to be exposed
EXPOSE 8000

# Run the app
CMD ["./pingeso-fusupo-backend-v2"]