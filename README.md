# Go gRPC Example

A simple gRPC client-server implementation in Go demonstrating a Calculator service.

## Project Structure

```
.
├── server/          # gRPC server implementation
│   ├── main.go
│   └── pb/          # Protocol buffer definitions and generated code
│       ├── calculator.proto
│       ├── calculator.pb.go
│       └── calculator_grpc.pb.go
└── client/          # gRPC client implementation
    ├── main.go
    └── pb/          # Protocol buffer definitions and generated code
        ├── calculator.proto
        ├── calculator.pb.go
        └── calculator_grpc.pb.go
```

## Features

- Calculator service with `Add` and `Subtract` RPC methods
- Protocol Buffers for service definition
- Client-server communication over TCP

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
   go mod download
   ```

2. Install dependencies for the client:
   ```bash
   cd client
   go mod download
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

The client performs an addition operation (1 + 1 by default). Modify the `num1` and `num2` variables in `client/main.go` to test with different values.

## Regenerating Protocol Buffers

If you modify the `.proto` file, regenerate the code:

```bash
cd server/pb
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    calculator.proto
```

Then copy the generated files to the client directory:

```bash
cp server/pb/calculator*.go client/pb/
```

## Service Definition

The service is defined in `calculator.proto`:

- **Service**: `Calculator`
- **Methods**: `Add`, `Subtract`
- **Request**: `CalcutorRequest` (contains `num1` and `num2` fields)
- **Response**: `CalcutorResponse` (contains a `result` field)

## License

This is a learning project.
