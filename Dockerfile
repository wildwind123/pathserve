FROM golang:latest as builder

# Install any necessary dependencies
RUN apt-get update && \
    apt-get install -y git

# Set the working directory to the project directory
WORKDIR /app

# Copy the source code into the container
COPY pathserve .
RUN mkdir bin
# Build the project for each target platform
RUN GOOS=linux GOARCH=amd64 go build -o bin/pathserve-amd64-linux
RUN GOOS=windows GOARCH=amd64 go build -o bin/pathserve-amd64-windows.exe
RUN GOOS=darwin GOARCH=amd64 go build -o bin/pathserve-amd64-mac


# Create a new stage to copy the built binaries to the host machine
FROM busybox:latest
COPY --from=builder /app /app

# Set the volume to copy the built binaries to the host machine
VOLUME ["/app"]

# Set the default command to do nothing, as this container only exists to share the volume
CMD ["tail", "-f", "/dev/null"]