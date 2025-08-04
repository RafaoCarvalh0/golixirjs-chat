# Matchmaker Service (Go)

This service is responsible for matching users into sessions. Built with Go for high performance and concurrency.

## Purpose
- Queue users
- Randomly pair users in a match
- Communicate the match with other services via messaging or API

## Running
This service can run independently. See the main.go or service documentation for details. 

```mermaid
sequenceDiagram
    participant User1
    participant Matchmaker
    participant User2

    User1->>Matchmaker: POST /create-match
    Matchmaker: waiting for a pair...

    User2->>Matchmaker: POST /create-match
    Matchmaker->>User1: match created
    Matchmaker->>User2: match created
``` 