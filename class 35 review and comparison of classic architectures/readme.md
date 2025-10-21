# Class 35: Review and Comparison of Classic Architectures

## Objectives

* Summarize and reinforce the understanding of **MVC**, **MVP**, **MVVM**, **Layered**, **Clean**, and **Hexagonal** architectures.
* Identify the strengths, weaknesses, and use cases of each architectural style.
* Learn how to select the most appropriate architecture for different project requirements.

---

## Overview

In this class, you’ll perform a comparative analysis of the most common **classic software architectures**. The goal is to understand how each architecture organizes code, manages dependencies, and handles separation of concerns.

---

## Key Architectures to Review

### 1. Model-View-Controller (MVC)

* **Strengths:** Clear separation between logic and UI, good for web apps.
* **Weaknesses:** Controller can become a bottleneck as complexity grows.
* **Use Case:** Web frameworks like Ruby on Rails, Django, and Gin.

### 2. Model-View-Presenter (MVP)

* **Strengths:** Improved testability over MVC, easier to mock Views.
* **Weaknesses:** More boilerplate code, Presenter can grow too large.
* **Use Case:** Desktop and mobile applications.

### 3. Model-View-ViewModel (MVVM)

* **Strengths:** Great for reactive UIs, clean data binding.
* **Weaknesses:** Complex for beginners, data binding can obscure flow.
* **Use Case:** UI frameworks (e.g., Android, React, WPF).

### 4. Layered Architecture

* **Strengths:** Simple and widely used; good organization by responsibility.
* **Weaknesses:** Can lead to tight coupling between layers.
* **Use Case:** Monolithic enterprise applications.

### 5. Clean Architecture

* **Strengths:** Emphasizes independence from frameworks and UI.
* **Weaknesses:** Higher initial complexity; requires strict discipline.
* **Use Case:** Large, long-term systems.

### 6. Hexagonal Architecture (Ports and Adapters)

* **Strengths:** Highly modular, easy to test and replace components.
* **Weaknesses:** May feel abstract for small teams or simple apps.
* **Use Case:** Systems requiring scalability and external integrations.

---

## Comparative Table

| Architecture | Key Focus                     | Dependency Direction   | Best For        | Difficulty |
| ------------ | ----------------------------- | ---------------------- | --------------- | ---------- |
| MVC          | Separation of logic and UI    | From UI to Model       | Web Apps        | Easy       |
| MVP          | Testability and UI control    | From View to Presenter | Desktop/Mobile  | Medium     |
| MVVM         | Reactive UIs                  | From View to ViewModel | Modern UIs      | Medium     |
| Layered      | Logical separation            | Between layers         | Enterprise Apps | Easy       |
| Clean        | Independence and testability  | Toward core            | Large systems   | Hard       |
| Hexagonal    | Decoupling and ports/adapters | Toward domain          | Modular systems | Hard       |

---

## Practical Exercise

1. Choose **two architectures** (e.g., MVC and Clean Architecture).
2. Implement a simple application (like a To-Do list or Book Manager) in both styles.
3. Write a short comparison document discussing:

   * Code organization
   * Ease of testing
   * Scalability and flexibility
   * Developer experience

---

## Additional Resources

* *Clean Architecture* by Robert C. Martin
* *Patterns of Enterprise Application Architecture* by Martin Fowler
* Online article: [Comparing MVC, MVP, and MVVM](https://www.geeksforgeeks.org/difference-between-mvc-mvp-and-mvvm-design-pattern/)

---

## Optional Challenge

Refactor one of your previous projects (like the Strategy or Observer example) to follow a **Clean** or **Hexagonal** architecture and document your process.

---

## Next Steps

➡️ **Next Class:** We will begin exploring **modern architectural paradigms**, starting with **Microservices Architecture**.
