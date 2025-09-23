# Class 33: Clean Architecture

## Objectives
- Understand the principles and goals of Clean Architecture.
- Learn how Clean Architecture separates concerns using layers.
- Implement a simple example in Go using Clean Architecture principles.
- Explore how Clean Architecture enhances testability, scalability, and maintainability.

## What is Clean Architecture?
Clean Architecture, proposed by Robert C. Martin (Uncle Bob), is a software design philosophy that emphasizes the separation of concerns. 
It ensures that business rules are independent of frameworks, databases, UI, or external services.

**Key ideas:**
- The inner layers (Entities, Use Cases) are pure and independent.
- Outer layers (Frameworks, UI, DB, APIs) depend on inner layers — not the other way around.
- Makes systems easier to test, extend, and maintain.

## UML Diagram

```
+---------------------+
|    Frameworks & UI  |
+---------------------+
          ↓
+---------------------+
|   Interface Adapters|
+---------------------+
          ↓
+---------------------+
|     Use Cases       |
+---------------------+
          ↓
+---------------------+
|      Entities       |
+---------------------+
```

## Conceptual Example
Imagine an **e-commerce system**:
- **Entities**: `Order`, `Product`, `User`
- **Use Cases**: `PlaceOrder`, `CancelOrder`
- **Interface Adapters**: Controllers, Presenters, Repositories
- **Frameworks & UI**: Web framework, Database, External APIs

The business rules (`PlaceOrder`) do not depend on the database or the UI. Instead, the infrastructure adapts to the business.

## Go Example

```go
package main

import "fmt"

// ===== Entities =====
type Order struct {
    ID     int
    Amount float64
}

// ===== Use Case =====
type OrderUseCase struct {
    repo OrderRepository
}

func (uc *OrderUseCase) PlaceOrder(order Order) {
    uc.repo.Save(order)
    fmt.Println("Order placed:", order.ID)
}

// ===== Repository Interface =====
type OrderRepository interface {
    Save(order Order)
}

// ===== Infrastructure Layer =====
type InMemoryOrderRepo struct {
    orders []Order
}

func (r *InMemoryOrderRepo) Save(order Order) {
    r.orders = append(r.orders, order)
    fmt.Println("Order saved to in-memory storage:", order.ID)
}

// ===== Main (Framework Layer) =====
func main() {
    repo := &InMemoryOrderRepo{}
    useCase := &OrderUseCase{repo: repo}

    order := Order{ID: 1, Amount: 100.0}
    useCase.PlaceOrder(order)
}
```

## Practical Example
Create a **task manager app** with the following layers:
1. **Entities**: `Task` (ID, Title, Status)
2. **Use Cases**: `AddTask`, `CompleteTask`
3. **Interface Adapters**: Console or HTTP handler
4. **Frameworks & DB**: Use an in-memory DB (later you can swap it for PostgreSQL without changing the use cases).

## Additional Resources
- Robert C. Martin, [Clean Architecture Book](https://www.oreilly.com/library/view/clean-architecture/9780134494272/)
- Blog: [Clean Architecture in Go](https://threedots.tech/post/introducing-clean-architecture/)
- Uncle Bob’s blog on [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

## Optional Challenge
Implement a **user authentication system** with Clean Architecture:
- **Entities**: User (ID, Email, PasswordHash)
- **Use Cases**: Register, Login
- **Adapters**: Repository interface + mock in-memory repository
- **Frameworks**: Simple CLI or HTTP server

## Next Steps
In the next class, we’ll explore **Hexagonal Architecture (Ports and Adapters)**, which shares similarities with Clean Architecture but emphasizes explicit boundaries and adapters for external systems.
