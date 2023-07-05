# Stage 1: build go
FROM golang:1.16 as build

# Copy the source files from the host
COPY . /src

# Set the working dir
WORKDIR /src

# build
RUN CGO_ENABLED=0 GOOS=linux go build -o kvs ./cmd/kvs

# Stage 2: build service
FROM scratch

# Copy the binaly from the build
COPY --from=build /src/kvs .

# Connect port to this Docker
EXPOSE 8080

# Run KVS
CMD ["/kvs"]