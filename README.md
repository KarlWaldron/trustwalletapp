# 🚀 Blockchain Client (Polygon JSON-RPC Proxy)

A simple Go-based HTTP API that proxies selected [Polygon](https://polygon-rpc.com) JSON-RPC methods, deployable on AWS ECS Fargate.

---

## 📌 Features

- Exposes two read-only endpoints:
  - `GET /block-number`: Fetch latest block number
  - `GET /block/{number}`: Fetch full block details
- Written in pure Go with `gorilla/mux`
- Unit tested using Go's standard library
- Lightweight Docker container
- Infrastructure-as-code via Terraform (ECS Fargate)

---

## ⚙️ How to Run Locally

### 1. Prerequisites

- Go 1.20+
- (optional) Docker
- (optional) Terraform & AWS CLI

### 2. Run the server

```bash
go run .
```

Server will start on port `8080`.

### 3. Test the API

```bash
curl http://localhost:8080/block-number
curl http://localhost:8080/block/20300000
```

---

## 🧪 Running Unit Tests

```bash
go test -v
```

Tests use `httptest` to mock Polygon RPC responses.

---

## 🐳 Docker Support

### 1. Build

```bash
docker build -t blockchain-client .
```

### 2. Run

```bash
docker run -p 8080:8080 blockchain-client
```

---

## ☁️ Terraform Deployment (AWS ECS Fargate)

### Prerequisites

- Docker image pushed to AWS ECR
- AWS credentials configured (via `aws configure` or environment variables)

### Steps

```bash
cd terraform
terraform init
terraform plan -var="container_image=<your-ecr-image-uri>"
terraform apply -var="container_image=<your-ecr-image-uri>"
```

Terraform will provision:

- VPC + Subnet
- ECS Cluster
- Security Group (port 8080 open)
- ECS Task Definition & Service

---

## 🛡 Production Readiness Suggestions

To make this production-grade:

- ✅ Add rate limiting and auth middleware (e.g. API key)
- ✅ Add retry and timeout logic in RPC calls
- ✅ Add logging (Zap, Logrus) and tracing
- ✅ Use Application Load Balancer with HTTPS termination
- ✅ Use parameterized secrets (e.g. AWS SSM or Secrets Manager)
- ✅ Add health check endpoints
- ✅ Implement Prometheus metrics
- ✅ Enable autoscaling (ECS Service Auto Scaling)

---


## 📁 Project Structure

```
.
├── main.go              # Entry point
├── handler.go           # HTTP handlers
├── client.go            # JSON-RPC logic
├── client_test.go       # Unit tests
├── Dockerfile
├── go.mod / go.sum
├── terraform/
│   ├── main.tf
│   ├── ecs.tf
│   ├── variables.tf
│   └── outputs.tf
└── README.md
```

---

## 🧑‍💻 Author

Built by Yunneng Wang. Contributions welcome!
