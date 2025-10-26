# Class 37: Event-Driven Architecture

## 🎯 Objectives
- Understand what Event-Driven Architecture (EDA) is and why it’s used.
- Learn the key components such as events, producers, consumers, and event brokers.
- Explore real-world use cases and how EDA improves scalability and decoupling.
- Implement a simple event-driven system in Go.

---

## 🧩 What is Event-Driven Architecture?
Event-Driven Architecture (EDA) is a design pattern that enables communication between services through **events** rather than direct calls.  
An *event* represents a change in state, such as “Order Created” or “User Registered.”

In this architecture:
- **Producers** emit events when something happens.
- **Consumers** listen and react to those events.
- **Event Brokers** (like Kafka, RabbitMQ, or NATS) route events between producers and consumers.

This allows for **loose coupling**, **asynchronous communication**, and **scalability**.

---

## 🧠 UML Diagram
Below is a conceptual representation of an Event-Driven Architecture:

```plaintext
+-------------+       +---------------+       +----------------+
|  Producer   | --->  |  Event Broker | --->  |   Consumer     |
| (OrderSvc)  |       |   (Kafka)     |       | (Email Service)|
+-------------+       +---------------+       +----------------+
         |                     |                      |
         |<---- Event Flow ---->|<---- Event Flow ---->|
```

---

## 💡 Conceptual Example
Let’s consider an **e-commerce system**.  
When a customer places an order, an event `OrderCreated` is published.  
Several services can independently respond to this event:

- **Inventory Service** reduces stock.  
- **Email Service** sends a confirmation email.  
- **Analytics Service** logs the order.

This happens **without direct dependencies** between these services.

---

## 🧰 Go Example
Below is a simple implementation using Go channels to simulate an Event Bus.

```go
package main

import "fmt"

// Event structure
type Event struct {
    Name string
    Data string
}

// EventBus simulates a basic event-driven system
type EventBus struct {
    subscribers map[string][]func(Event)
}

func NewEventBus() *EventBus {
    return &EventBus{subscribers: make(map[string][]func(Event))}
}

func (b *EventBus) Subscribe(eventName string, handler func(Event)) {
    b.subscribers[eventName] = append(b.subscribers[eventName], handler)
}

func (b *EventBus) Publish(event Event) {
    if handlers, ok := b.subscribers[event.Name]; ok {
        for _, handler := range handlers {
            handler(event)
        }
    }
}

func main() {
    bus := NewEventBus()

    // Subscribers
    bus.Subscribe("OrderCreated", func(e Event) {
        fmt.Println("Inventory Service:", e.Data, "-> Decreasing stock")
    })
    bus.Subscribe("OrderCreated", func(e Event) {
        fmt.Println("Email Service:", e.Data, "-> Sending confirmation email")
    })

    // Publish Event
    bus.Publish(Event{Name: "OrderCreated", Data: "Order #123"})
}
```

---

## 🧪 Practical Example
Extend the Go implementation by:
1. Adding a new event `PaymentProcessed`.
2. Creating a service that listens to it (e.g., “Shipping Service”).
3. Publishing multiple events to simulate an order workflow.

---

## 📚 Additional Resources
- [Martin Fowler – Event-Driven Architecture](https://martinfowler.com/articles/201701-event-driven.html)
- [AWS Docs – Event-driven architecture](https://aws.amazon.com/event-driven-architecture/)
- [Apache Kafka Documentation](https://kafka.apache.org/documentation/)

---

## 💥 Optional Challenge
Refactor your event-driven system to include **asynchronous event handling** using goroutines.  
Then, simulate **multiple concurrent producers** publishing different events at the same time.

---

## 🚀 Next Steps
Next class (Class 38): **Serverless Architecture** – Learn how to design systems that respond to events without managing servers.
