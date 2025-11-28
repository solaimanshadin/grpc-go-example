# Go gRPC Example

A simple gRPC client-server implementation in Go demonstrating a basic "Hello World" service.

## Project Structure

```
.
├── server/          # gRPC server implementation
│   ├── main.go
│   └── pb/          # Protocol buffer definitions and generated code
│       ├── helloworld.proto
│       ├── helloworld.pb.go
│       └── helloworld_grpc.pb.go
└── client/          # gRPC client implementation
    ├── main.go
    └── pb/          # Protocol buffer definitions and generated code
        ├── helloworld.proto
        ├── helloworld.pb.go
        └── helloworld_grpc.pb.go
```

## Features

- Simple gRPC service with a `SayHello` RPC method
- Protocol Buffers for service definition


## Prerequisites

- Go 1.16 or higher
- Protocol Buffer compiler (`protoc`)
- Go plugins for Protocol Buffers:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## Setup

1. Install dependencies for the server:
   ```bash
   cd server
   go mod tidy
   ```

2. Install dependencies for the client:
   ```bash
   cd client
   go mod tidy
   ```

## Running the Application

### Start the Server

```bash
cd server
go run main.go
```

The server will start listening on `localhost:50051`.

### Run the Client

In a separate terminal:

```bash
cd client
go run main.go
```

By default, the client sends a request with the name "world". You can specify a custom name:

```bash
go run main.go -name="YourName"
```

## Regenerating Protocol Buffers

If you modify the `.proto` file, regenerate the code:

```bash
cd server/pb
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```

Then copy the generated files to the client directory:

```bash
cp server/pb/helloworld*.go client/pb/
```



