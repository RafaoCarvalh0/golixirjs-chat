# Matchmaker Service (Go)

This service is responsible for matching users into chat sessions. Built with Go for high performance and concurrency.

## Purpose
- Queue online users
- Randomly pair users for chat
- Communicate with other services via messaging or API

## Running
This service can be run independently. See the main.go or service documentation for details. 

## Exemplo de Fluxo em Diagrama

```mermaid
sequenceDiagram
    participant User1
    participant Matchmaker
    participant User2

    User1->>Matchmaker: POST /queue
    Matchmaker->>User1: Aguardando par...

    User2->>Matchmaker: POST /queue
    Matchmaker->>User1: Match encontrado! (sala 123)
    Matchmaker->>User2: Match encontrado! (sala 123)
``` 