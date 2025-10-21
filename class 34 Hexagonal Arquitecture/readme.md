# Class 34: Hexagonal Architecture

## Objectives
- Understand the principles of Hexagonal Architecture (also known as Ports and Adapters).
- Learn how it helps to decouple business logic from infrastructure.
- Explore how Hexagonal Architecture compares to traditional layered architecture.
- Implement a simple example in Go following the ports and adapters approach.

---

## What is Hexagonal Architecture?
Hexagonal Architecture, proposed by Alistair Cockburn, is a software design pattern that aims to isolate the core business logic (the "domain") from external concerns such as databases, APIs, user interfaces, or frameworks.  

It is structured around **ports** (interfaces that define how the application communicates) and **adapters** (implementations that connect external systems to the ports).

Key benefits:
- High testability of business logic.
- Independence from frameworks or external tools.
- Flexibility to replace infrastructure components without changing the core.

---

## UML Diagram

```
        +----------------------+
        |      Application     |
        |      (Domain)        |
        +----------------------+
             ^            ^
             |            |
      +------+            +------+
      | Ports |            | Ports |
      +------+            +------+
         ^                   ^
         |                   |
   +-----------+       +-------------+
   | Adapters  |       |  Adapters   |
   | (DB, UI)  |       | (API, CLI)  |
   +-----------+       +-------------+
```

---

## Conceptual Example
Imagine a **bank account application**:
- The **core domain** contains the business rules: deposit, withdraw, check balance.
- **Ports** define interfaces (e.g., `AccountRepository`).
- **Adapters** implement these ports:
  - One adapter connects to a database.
  - Another adapter connects to a REST API.

This allows replacing a database (e.g., PostgreSQL → MongoDB) or interface (e.g., REST → CLI) without changing the business logic.

---

## Go Example

```go
// Domain (Core)
package domain

type Account struct {
    ID      string
    Balance float64
}

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
    if amount > a.Balance {
        return false
    }
    a.Balance -= amount
    return true
}

// Port
type AccountRepository interface {
    Save(account *Account) error
    FindByID(id string) (*Account, error)
}
```

```go
// Adapter: In-memory repository
package adapter

import (
    "errors"
    "hexagonal/domain"
)

type InMemoryRepo struct {
    accounts map[string]*domain.Account
}

func NewInMemoryRepo() *InMemoryRepo {
    return &InMemoryRepo{accounts: make(map[string]*domain.Account)}
}

func (r *InMemoryRepo) Save(account *domain.Account) error {
    r.accounts[account.ID] = account
    return nil
}

func (r *InMemoryRepo) FindByID(id string) (*domain.Account, error) {
    acc, exists := r.accounts[id]
    if !exists {
        return nil, errors.New("account not found")
    }
    return acc, nil
}
```

```go
// Application Service (uses port)
package service

import "hexagonal/domain"

type AccountService struct {
    repo domain.AccountRepository
}

func NewAccountService(r domain.AccountRepository) *AccountService {
    return &AccountService{repo: r}
}

func (s *AccountService) Deposit(id string, amount float64) error {
    acc, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }
    acc.Deposit(amount)
    return s.repo.Save(acc)
}
```

---

## Practical Example
- Implement a **task manager** app:
  - Core domain: `Task` entity and rules (create, mark as done).
  - Ports: `TaskRepository` interface.
  - Adapters:
    - In-memory repository.
    - JSON file repository.
    - REST API handler.
- Swap between in-memory and JSON persistence without touching domain logic.

---

## Additional Resources
- Alistair Cockburn: [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)  
- Book: *Clean Architecture* by Robert C. Martin.  
- Blog: [Hexagonal Architecture explained](https://blog.ndepend.com/hexagonal-architecture/)  

---

## Optional Challenge
- Extend the **bank account example** to include:
  - Multiple adapters (REST API + CLI).
  - Database-backed repository.
- Demonstrate replacing one adapter with another without changing the domain.

---

## Next Steps
In the next class, we will explore **Event-Driven Architecture (EDA)**, a style where systems communicate through events and message brokers.
