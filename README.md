# CRUD App

A full-stack web application featuring a **NextJS 15** frontend and **Golang** backend implementing **Hexagonal Architecture** (Ports & Adapters pattern).

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-15.3-black?style=flat&logo=next.js)](https://nextjs.org/)
[![React](https://img.shields.io/badge/React-19-61DAFB?style=flat&logo=react)](https://react.dev/)
[![TailwindCSS](https://img.shields.io/badge/TailwindCSS-4.1-38B2AC?style=flat&logo=tailwind-css)](https://tailwindcss.com/)
[![DevContainer](https://img.shields.io/badge/DevContainer-Ready-blue?style=flat&logo=docker)](https://containers.dev/)

---

## 🏗️ Architecture Overview

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              CRUD Application                               │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                        Frontend (ui/)                                │   │
│  │  NextJS 15 • React 19 • TypeScript • TailwindCSS 4 • Cypress        │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                   │                                         │
│                                   ▼ HTTP/REST                               │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │                        Backend (api/)                                │   │
│  │  ┌─────────────────┐  ┌─────────────────┐  ┌─────────────────────┐  │   │
│  │  │   Left Adapters │  │   Application   │  │   Right Adapters    │  │   │
│  │  │   (HTTP/gRPC)   │◄►│   Core/Domain   │◄►│   (PostgreSQL)      │  │   │
│  │  └─────────────────┘  └─────────────────┘  └─────────────────────┘  │   │
│  │        Gin + gRPC           Hexagonal            Squirrel SQL        │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📁 Project Structure

```
crud-app/
├── .devcontainer/          # DevContainer configuration
│   ├── devcontainer.json   # VS Code DevContainer settings
│   ├── docker-compose.yml  # Multi-container Docker setup
│   ├── Dockerfile          # Go 1.24 base image
│   └── postCreateCommand.sh
│
├── api/                    # Backend - Golang Hexagonal Architecture
│   ├── cmd/                # Application entry points
│   │   ├── app/            # Main HTTP server entry
│   │   ├── appointment/    # Appointment service
│   │   ├── company/        # Company service
│   │   ├── patient/        # Patient service
│   │   └── ...             # Other domain services
│   │
│   └── internal/           # Private application code
│       ├── adapters/       # Framework adapters
│       │   └── framework/
│       │       ├── left/   # Driving adapters (HTTP, gRPC)
│       │       └── right/  # Driven adapters (Database)
│       ├── application/    # Business logic
│       │   ├── api/        # Application services
│       │   └── core/       # Domain models
│       │       └── domain/ # Entities & value objects
│       └── ports/          # Interface definitions
│           ├── left.go     # Driving ports (API)
│           └── right.go    # Driven ports (DB)
│
├── ui/                     # Frontend - NextJS 15
│   ├── app/                # App Router pages
│   │   ├── page.tsx        # Home page
│   │   ├── layout.tsx      # Root layout
│   │   └── globals.css     # Global styles
│   ├── public/             # Static assets
│   ├── package.json        # NPM dependencies
│   └── tsconfig.json       # TypeScript config
│
├── .github/                # GitHub configuration
├── .vscode/                # VS Code settings
├── LICENSE                 # Project license
└── README.md               # This file
```

---

## 🛠️ Technology Stack

### Frontend (`ui/`)

| Technology      | Version | Purpose                         |
| --------------- | ------- | ------------------------------- |
| **Next.js**     | 15.3.2  | React framework with App Router |
| **React**       | 19.0.0  | UI component library            |
| **TypeScript**  | 5.x     | Type-safe JavaScript            |
| **TailwindCSS** | 4.1.7   | Utility-first CSS framework     |
| **Cypress**     | 14.4.0  | End-to-end testing              |
| **ESLint**      | 9.x     | Code linting                    |

### Backend (`api/`)

| Technology      | Version | Purpose              |
| --------------- | ------- | -------------------- |
| **Go**          | 1.23+   | Programming language |
| **Gin**         | 1.10.0  | HTTP web framework   |
| **gRPC**        | 1.72.1  | High-performance RPC |
| **Protobuf**    | 1.36.6  | Protocol buffers     |
| **Squirrel**    | 1.5.4   | SQL query builder    |
| **lib/pq**      | 1.10.9  | PostgreSQL driver    |
| **google/uuid** | 1.6.0   | UUID generation      |

### DevContainer Features

- **Go 1.24** (Debian Bookworm)
- **Node.js LTS** with npm, yarn, pnpm
- **PostgreSQL Client** v13
- **AWS CLI** & LocalStack
- **Dapr CLI** for microservices
- **Protoc** compiler for gRPC
- **Cypress & Playwright** for testing

---

## 🚀 Getting Started

### Prerequisites

- [Docker Desktop](https://www.docker.com/products/docker-desktop/)
- [VS Code](https://code.visualstudio.com/)
- [Dev Containers Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

### Quick Start

1. **Clone the repository**

   ```bash
   git clone <repository-url>
   cd crud-app
   ```

2. **Open in VS Code**

   ```bash
   code .
   ```

3. **Reopen in Container**
   - Press `F1` and select **"Dev Containers: Reopen in Container"**
   - Wait for the container to build (first time takes a few minutes)

4. **Start the Backend**

   ```bash
   cd api
   go run cmd/app/main.go
   ```

   Server runs on `http://localhost:8080`

5. **Start the Frontend**
   ```bash
   cd ui
   npm run dev
   ```
   App runs on `http://localhost:3000`

---

## 📡 API Endpoints

| Method   | Endpoint       | Description        | Status         |
| -------- | -------------- | ------------------ | -------------- |
| `GET`    | `/`            | Welcome message    | ✅ Implemented |
| `GET`    | `/ping`        | Health check       | ✅ Implemented |
| `GET`    | `/companies`   | List all companies | 🚧 Pending     |
| `GET`    | `/company/:id` | Get company by ID  | 🚧 Pending     |
| `POST`   | `/company`     | Create new company | ✅ Implemented |
| `PATCH`  | `/company/:id` | Update company     | 🚧 Pending     |
| `DELETE` | `/company/:id` | Delete company     | 🚧 Pending     |

---

## 🧪 Development Scripts

### Frontend (`ui/`)

```bash
npm run dev      # Start development server with Turbopack
npm run build    # Create production build
npm run start    # Start production server
npm run lint     # Run ESLint
```

### Backend (`api/`)

```bash
go run cmd/app/main.go       # Run HTTP server
go test ./...                 # Run all tests
go build -o bin/api cmd/app/main.go  # Build binary
```

### gRPC Development

```bash
# Generate Go code from proto files
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

## 🔧 DevContainer Configuration

The project includes a comprehensive DevContainer setup:

- **Container Name**: `crud-app`
- **Resource Limits**: 8GB RAM, 4 CPUs (reserved: 2GB RAM, 1 CPU)
- **Forwarded Ports**: 5432 (PostgreSQL), 5672 (RabbitMQ), 6379 (Redis), 8080 (API), 9000, 15672
- **User**: `vscode` (non-root for security)
- **Network**: `keycloak-dbs-brokers_backend_network`

---

## 📚 Related Documentation

- [ANALYSIS.md](./ANALYSIS.md) - Detailed technical analysis of the codebase
- [api/README.md](./api/README.md) - Backend-specific documentation
- [ui/README.md](./ui/README.md) - Frontend-specific documentation

---

## 📄 License

This project is licensed under the terms specified in the [LICENSE](./LICENSE) file.
