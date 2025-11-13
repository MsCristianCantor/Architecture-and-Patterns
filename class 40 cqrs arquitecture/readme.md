# Class 40: CQRS Architecture

## 🎯 Objectives
- Understand the Command Query Responsibility Segregation (CQRS) pattern.
- Learn when and why to apply CQRS in software design.
- Implement a simple CQRS example using Go.
- Compare CQRS with traditional CRUD architectures.

---

## 🧠 What is CQRS?
CQRS stands for **Command Query Responsibility Segregation**, a software architecture pattern that separates **read** and **write** operations into different models.

In traditional architectures, the same model is used to perform all CRUD (Create, Read, Update, Delete) operations. In CQRS, commands and queries are **handled separately**, allowing for better scalability, performance, and maintainability.

- **Command**: performs an action that changes the system state.
- **Query**: retrieves data without modifying the system state.

This separation enables:
- Different optimization strategies for reads and writes.
- Easier scalability for read-heavy or write-heavy systems.
- Potential for eventual consistency and event-driven updates.

---

## 📊 UML Diagram

```
+----------------------+          +----------------------+
|   Command Handler    |          |     Query Handler    |
|----------------------|          |----------------------|
| + Execute(command)   |          | + Handle(query)      |
+----------+-----------+          +----------+-----------+
           |                                 |
           v                                 v
    +--------------+                  +--------------+
    |  Write Model |                  |   Read Model |
    +--------------+                  +--------------+
           |                                 |
           v                                 v
      +---------+                       +---------+
      | Storage |                       | Storage |
      +---------+                       +---------+
```

---

## 💡 Conceptual Example

Imagine an **e-commerce system**:

- When a user places an order → it triggers a **Command** (`PlaceOrderCommand`) handled by a **Command Handler** that updates the write model.
- When a user views their order history → it triggers a **Query** (`GetOrdersQuery`) handled by a **Query Handler** that reads from the read model (possibly a denormalized database).

This allows scaling reads (queries) separately from writes (commands).

---

## 🧩 Go Example

```go
package main

import "fmt"

// Command
type CreateUserCommand struct {
    Name string
    Age  int
}

// Command Handler
type CommandHandler struct {
    users map[string]int
}

func (h *CommandHandler) Handle(cmd CreateUserCommand) {
    h.users[cmd.Name] = cmd.Age
    fmt.Printf("User %s created, age %d\n", cmd.Name, cmd.Age)
}

// Query
type GetUserQuery struct {
    Name string
}

// Query Handler
type QueryHandler struct {
    users map[string]int
}

func (h *QueryHandler) Handle(query GetUserQuery) (int, bool) {
    age, exists := h.users[query.Name]
    return age, exists
}

func main() {
    commandHandler := &CommandHandler{users: make(map[string]int)}
    queryHandler := &QueryHandler{users: commandHandler.users}

    commandHandler.Handle(CreateUserCommand{Name: "Alice", Age: 30})

    age, found := queryHandler.Handle(GetUserQuery{Name: "Alice"})
    if found {
        fmt.Printf("Found user Alice with age %d\n", age)
    }
}
```

---

## 🧠 Practical Example

You can build a **Go web API** using CQRS principles:
- Create separate packages for `commands` and `queries`.
- Use different databases or tables for reading and writing.
- Optionally integrate an **event bus** to synchronize read models asynchronously.

For instance:
- `POST /users` → handled by a command that writes data.
- `GET /users` → handled by a query optimized for reading.

---

## 📚 Additional Resources
- [Microsoft Docs – CQRS Pattern](https://learn.microsoft.com/en-us/azure/architecture/patterns/cqrs)
- [Martin Fowler – CQRS](https://martinfowler.com/bliki/CQRS.html)
- [Greg Young – Original CQRS Talk](https://cqrs.wordpress.com/)

---

## 🧩 Optional Challenge
Implement a CQRS-based **task manager** in Go:
- Commands: CreateTask, CompleteTask.
- Queries: GetTasks, GetCompletedTasks.
- Store data in memory or use a lightweight database like SQLite.

---

## 🚀 Next Steps
Next class: **Event Sourcing** — understanding how to persist and rebuild system state using events.
