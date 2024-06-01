# Use the official Golang image as a build stage
FROM golang:1.21 as build

WORKDIR /usr/src/app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download && go mod tidy && go mod verify

# Copy the entire project
COPY . .

# Set working directory for the build
WORKDIR /usr/src/app/aggregator/cmd

# Build the Go app and capture errors
RUN go build -v -o /usr/local/bin/aggregator ./... || { echo 'Go build failed'; exit 1; }

# Use a minimal image for the final stage
FROM debian:latest

# Copy the built binary from the build stage
COPY --from=build /usr/local/bin/aggregator /usr/local/bin/aggregator

# Set the entrypoint and default command
ENTRYPOINT ["aggregator"]
CMD ["--config=config-files/aggregator-docker-compose.yaml", "--credible-squaring-deployment=contracts/script/output/31337/credible_squaring_avs_deployment_output.json", "--ecdsa-private-key=0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"]
