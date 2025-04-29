# CRM Microservices project

A modular CRM system built using **microservices architecture**, **gRPC**, **Go** and **Docker**.

## Project Structure

crm-golang/
├── api-gateway/
│   ├── go.mod
│   ├── main.go
│   ├── grpc-clients/      # здесь генерируем pb.go, pb_grpc.go для клиента
│   └── handlers/
│       └── user_handler.go
│
├── item-service/
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── grpc/     
│   │   │   ├── item-grpc/
│   │   │   │   ├── item.pb_grpc.go
│   │   │   │   ├── item.pb.go
│   │   │   │   └── item.proto
│   │   │   └── models/
│   │   │       ├── structures.pb.go
│   │   │       └── structures.proto
│   │   └── models/
│   │       └── item.go
│   ├── go.sum
│   └── go.mod
│
├── lead-service/
│   ├── cmd
│   │   └── main.go
│   ├── internal
│   │   ├── grpc/     
│   │   │   ├── lead-grpc/
│   │   │   │   ├── lead.pb_grpc.go
│   │   │   │   ├── lead.pb.go
│   │   │   │   └── lead.proto
│   │   │   └── models/
│   │   │       ├── structures.pb.go
│   │   │       └── structures.proto
│   │   └── models/
│   │       └── lead.go
│   ├── go.sum
│   └── go.mod
├── .dockerignore
└── .gitignore             
