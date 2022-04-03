# Stage 1: Compile the binary

FROM golang:1.16 as build

# Copy source file from host
COPY . /src

# Set the work directory
WORKDIR /src

# Build the binary!
RUN CGO_ENABLED=0 GOOS=linux go build -o pet-adoption-server

# Stage 2: Build the application image

FROM alpine:latest

# Copy binary from build container
COPY --from=build /src/pet-adoption-server .

# Todo: later add pem keys for https

EXPOSE 8080

# Tell docker what command to use on docker run
CMD ["/pet-adoption-server"]

# docker build -t pet-adoption-server:latest .
# docker run -it -p 8080:8080 pet-adoption-server