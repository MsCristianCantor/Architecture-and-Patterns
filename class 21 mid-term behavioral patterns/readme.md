# Class 21 – Mid-term Review + Behavioral Patterns Exercises

## 🎯 Objectives
- Review and reinforce understanding of **Chain of Responsibility**, **Command**, **Iterator**, **Mediator**, **Memento**, and **Observer**.
- Identify the problem each pattern solves and the scenarios where it is most effective.
- Practice implementing and combining these patterns in Go.
- Strengthen pattern recognition through real-world inspired exercises.

---

## 📖 What is this session about?
This is a **focused review** of the first set of behavioral design patterns we’ve learned.  
We will:
1. Revisit the **intent**, **structure**, and **common use cases** of each pattern.
2. Analyze small UML diagrams for each pattern.
3. Implement a combined exercise involving at least two patterns.
4. Discuss trade-offs and when *not* to use a pattern.

---

## 🗂️ Quick UML Recap

### Chain of Responsibility

[Handler] → [Handler] → [Handler]

Each handler decides to process a request or pass it to the next handler.

### Command

[Invoker] → [Command Interface] → [Concrete Command] → [Receiver]

Encapsulates an action as an object.

### Iterator

[Collection] → [Iterator] → [Element]

Provides a way to traverse elements without exposing the underlying structure.

### Mediator

[Colleague 1] ↔
[Mediator]
[Colleague 2] ↔

Centralizes complex communications between objects.

### Memento

[Originator] ↔ [Memento] → [Caretaker]

Captures and restores an object’s state without exposing internal details.

### Observer

[Subject] → [Observer 1]
→ [Observer 2]

Notifies multiple observers of state changes.

---

## 💡 Conceptual Example (Combined Use)
**Scenario:** A support ticket system.
- **Chain of Responsibility** → Route tickets based on priority or department.
- **Observer** → Notify support agents when a ticket is updated.
- **Memento** → Save ticket history for rollback.

---

## 🖥 Example in Go (Chain of Responsibility + Observer)

```go
package main

import "fmt"

// Chain of Responsibility
type Handler interface {
    SetNext(Handler)
    HandleRequest(level int)
}

type BaseHandler struct{ next Handler }
func (b *BaseHandler) SetNext(next Handler) { b.next = next }
func (b *BaseHandler) HandleRequest(level int) {
    if b.next != nil {
        b.next.HandleRequest(level)
    }
}

type Level1Handler struct{ BaseHandler }
func (l *Level1Handler) HandleRequest(level int) {
    if level == 1 {
        fmt.Println("Level 1 handled the request.")
        notifyObservers("Ticket handled at Level 1")
    } else {
        l.BaseHandler.HandleRequest(level)
    }
}

type Level2Handler struct{ BaseHandler }
func (l *Level2Handler) HandleRequest(level int) {
    if level == 2 {
        fmt.Println("Level 2 handled the request.")
        notifyObservers("Ticket handled at Level 2")
    } else {
        l.BaseHandler.HandleRequest(level)
    }
}

// Observer
type Observer interface {
    Update(msg string)
}

type Agent struct{ name string }
func (a *Agent) Update(msg string) {
    fmt.Printf("[%s] received notification: %s\n", a.name, msg)
}

var observers []Observer

func addObserver(o Observer) { observers = append(observers, o) }
func notifyObservers(msg string) {
    for _, o := range observers {
        o.Update(msg)
    }
}

func main() {
    // Setup observers
    addObserver(&Agent{name: "Alice"})
    addObserver(&Agent{name: "Bob"})

    // Setup chain
    h1 := &Level1Handler{}
    h2 := &Level2Handler{}
    h1.SetNext(h2)

    h1.HandleRequest(1)
    h1.HandleRequest(2)
}
```

## 🏋️ Practical Exercise

Task:

1. Choose **two or more** of the reviewed patterns.

2. Build a **task management system** where:

    * **Chain of Responsibility** routes tasks to the right department.

    * **Command** encapsulates task actions (assign, complete, delete).

    * **Iterator** allows listing tasks without exposing internal structure.

    * **Mediator** coordinates communications between different modules.

    * **Memento** stores snapshots of a task for rollback.

    * **Observer** notifies subscribers of task changes.

---

## 📚 Additional Resources

* Refactoring Guru – [Behavioral Patterns](https://refactoring.guru/design-patterns/behavioral-patterns)

* Go patterns reference: [github.com/tmrts/go-patterns](https://github.com/tmrts/go-patterns)

* Head First Design Patterns” – Chapters on behavioral patterns.

---

## 🎯 Optional Challenge

* Implement a **multi-pattern solution** combining **at least 3** of the reviewed patterns.

* Justify your pattern selection in a short README.

---

## 🔜 What’s Next?

In the next class, we’ll continue with (Class 26).