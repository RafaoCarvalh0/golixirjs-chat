# 📚 Project Overview: RandomMatch Chat System

## 🎯 Core Idea

This system is a distributed chat platform that connects two random online users for real-time communication. If no other user is available, the system will eventually fall back to a simulated AI chat agent. For now, we'll focus on the real-time peer-to-peer feature.

The goal is to simulate a lightweight, anonymous social interaction tool — like "Chatroulette", but in text, and backed by modern backend engineering concepts.

---

## 🧠 Learning Goals

This project aims to explore and practice key backend development concepts:

- **Hexagonal Architecture (Ports & Adapters)**: To enforce separation of concerns, making each microservice modular and testable.
- **Microservices Architecture**: Three independent services communicating via REST or messaging, written in different languages.
- **Asynchronous Communication**: Using Redis Pub/Sub or a message broker (e.g., NATS) for decoupled message passing.
- **Service Coordination**: Implementing real-time match logic and fallback behavior.
- **Deployment & CI/CD**: Unified monorepo setup for containerized multi-service deploys.

---

## 🛠️ Technologies and Service Roles

| Service | Language | Role | Tech Stack | Reason |
|--------|----------|------|------------|--------|
| **Matchmaking Service** | Go (Gin) | Handles queueing of online users and random pairing | Gin, Redis Pub/Sub | Go provides great performance and concurrency (via goroutines), ideal for fast matching logic |
| **Chat Service** | Elixir (Phoenix) | Manages live chat sessions between users | Phoenix Channels, ETS/Redis | Phoenix excels in real-time communication via WebSockets, with scalability built-in |
| **User Management/API Gateway** | Node.js (NestJS) | Handles user registration, session state, and exposes public endpoints | NestJS, JWT/Auth, REST | Nest provides structure and DX for building modular APIs quickly |

---

## 📦 Monorepo Structure

```bash
.
├── micro_services/
│   ├── matchmaking_service/  # Go (Gin)
│   ├── chat_service/         # Elixir (Phoenix)
│   └── user_api_gateway/     # Node (NestJS)
├── docker/
│   └── docker-compose.yml    # Multi-service orchestration
├── README.md
└── .github/
    └── workflows/            # CI/CD pipelines
```