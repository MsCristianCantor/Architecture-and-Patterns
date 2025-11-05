# Class 39: Service-Oriented Architecture (SOA)

## 🎯 Objectives
- Understand the core principles of Service-Oriented Architecture (SOA).
- Identify how SOA differs from microservices and monolithic architectures.
- Learn how to design modular, reusable, and interoperable services.
- Implement a simple SOA example in Go.

## 🧠 What is Service-Oriented Architecture?
Service-Oriented Architecture (SOA) is a software architectural pattern where software components provide services to other components over a network through a communication protocol. These services are designed to be **loosely coupled**, **reusable**, and **interoperable**, enabling distributed system integration.

### Key Characteristics
- **Loose coupling:** Services interact through well-defined interfaces.
- **Reusability:** Services can be used across different applications.
- **Interoperability:** Services communicate regardless of the underlying platform.
- **Discoverability:** Services can be found and invoked dynamically.

SOA was a foundational concept that influenced **microservices architecture**, but it typically uses **enterprise service buses (ESBs)** or centralized communication layers.

## 🧩 UML Diagram
A simplified UML representation of SOA:

```
+----------------------+       +----------------------+
|     Client App       | --->  |     Service Bus      |
+----------------------+       +----------+-----------+
                                         |
                                         v
                        +-------------------------------+
                        |      Auth Service (SOAP/REST) |
                        +-------------------------------+
                        |     Billing Service (SOAP)    |
                        +-------------------------------+
                        |     Inventory Service (REST)  |
                        +-------------------------------+
```

## 💡 Conceptual Example
Imagine an **e-commerce system** where different services handle authentication, billing, and inventory management. Each service is independent and communicates through a service bus or message broker.

## 🧰 Go Example
A simplified SOA-style interaction using REST services:

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Product struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Stock int    `json:"stock"`
}

// Inventory Service
func inventoryHandler(w http.ResponseWriter, r *http.Request) {
    products := []Product{
        {ID: "1", Name: "Hoodie", Stock: 20},
        {ID: "2", Name: "Sweater", Stock: 10},
    }
    json.NewEncoder(w).Encode(products)
}

// Client that consumes the service
func main() {
    http.HandleFunc("/inventory", inventoryHandler)
    fmt.Println("Inventory Service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

This example represents one **service** in a larger SOA system. Other services (like billing or authentication) could be implemented similarly.

## 🧪 Practical Example
Create a **mini SOA project** in Go with the following services:
1. **Auth Service** — handles login and authentication.
2. **Inventory Service** — manages product stock.
3. **Order Service** — coordinates order processing by calling other services.

Use HTTP to communicate between services, simulating a service bus or API gateway.

## 📚 Additional Resources
- [Service-Oriented Architecture Explained – IBM](https://www.ibm.com/cloud/learn/soa)
- [SOA vs Microservices – Red Hat](https://www.redhat.com/en/topics/microservices/what-is-service-oriented-architecture)
- Book: *Enterprise SOA: Service-Oriented Architecture Best Practices* by Dirk Krafzig et al.

## 💪 Optional Challenge
Implement a **service discovery mechanism** (for example, using a simple registry or Consul) to dynamically connect your Go services.

## 🚀 Next Steps
Next class, we’ll explore **API Gateway patterns** and how they connect with both SOA and microservices for unified service management.
