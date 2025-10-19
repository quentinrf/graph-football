# Architecture Overview

## Overview

This service is built using the **Hexagonal Architecture (Ports and Adapters)** pattern.

The goal is to keep the **domain logic** (graph modeling and traversal) completely independent from infrastructure concerns like AWS, databases, or APIs.  
This approach makes the codebase modular, testable, and easy to extend with new storage backends or services.

---

## Key Concepts

### 1. Domain-Centric Design

At the center of the architecture is the **domain** — pure Go code that models graphs and their operations.

Everything outside the domain interacts with it through well-defined **ports** (interfaces).  
External systems (AWS, Salesforce, Snowflake, etc.) are integrated through **adapters** that implement those interfaces.

### 2. Layers

| Layer | Description | Dependencies |
|-------|--------------|---------------|
| **Domain** | Core business logic (graph structure, traversal algorithms, domain interfaces). | None |
| **Application** | Use cases that orchestrate domain operations (e.g., creating, saving, searching graphs). | Domain |
| **Adapters** | Infrastructure code implementing domain interfaces (e.g., DynamoDB, in-memory). | External SDKs |
| **Infrastructure / Entry Point** | Wiring, dependency injection, configuration, HTTP servers. | All other layers |

---

## Summar Diagram

```pgsql
          +---------------------------+
          |       Application Layer   |
          |  (Use Cases, Orchestration)|
          +-------------+-------------+
                        |
         [Ports]   GraphRepository Interface
                        |
          +-------------+-------------+
          |   Domain Layer (Core)     |
          | Graph entities, logic     |
          +-------------+-------------+
                        |
       [Adapters] DynamoDB | Memory | Future integrations
                        |
          +-------------+-------------+
          | Infrastructure & main.go  |
          | (Configuration & Wiring)  |
          +---------------------------+

```

## Benefits

Key Benefits
- ✅ Testability – Domain and application logic can be tested with an in-memory repository.
- ✅ Replaceable infrastructure – Switch DynamoDB for another store with no domain changes.
- ✅ Modular design – Each layer has clear, single-purpose responsibility.
- ✅ Extendable – Easily add adapters (e.g., Snowflake projections, Salesforce syncs) without touching core logic.
- ✅ Idiomatic Go – Clean, flat packages; internal/ enforces encapsulation.