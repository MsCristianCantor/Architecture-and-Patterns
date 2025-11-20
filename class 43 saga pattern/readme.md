# Class 43: Saga Pattern

## 🎯 Objectives

-   Understand the Saga pattern and why it is used in distributed
    systems.
-   Learn the differences between **choreography** and
    **orchestration**-based sagas.
-   Model long-running, multi-step transactions with compensating
    actions.
-   Apply the Saga pattern to a realistic business workflow.

------------------------------------------------------------------------

## 📘 What Is the Saga Pattern?

The **Saga Pattern** is a design pattern used in **distributed systems**
to manage data consistency **without using distributed transactions**.

A Saga is a **sequence of steps (local transactions)**.\
If one step fails, the system triggers **compensating actions** to undo
the previous steps.

------------------------------------------------------------------------

## 🧩 Why Sagas Exist

Traditional ACID transactions are not practical across multiple services
in a microservice architecture.\
The Saga pattern allows you to maintain consistency **eventually**,
while keeping each service independent.

------------------------------------------------------------------------

## 🧩 Saga Execution Models

### 1. **Choreography (Event-Driven Sagas)**

Each service: - Performs its local transaction\
- Emits an event\
- Another service reacts to that event\
- No central coordinator

**Pros:** - Simple\
- Decentralized\
- Loose coupling

**Cons:** - Hard to trace\
- Business logic is distributed\
- Can become "event spaghetti"

------------------------------------------------------------------------

### 2. **Orchestration (Central Controller)**

A **Saga orchestrator**: - Tells each service what to do\
- Waits for responses\
- Triggers compensations if needed

**Pros:** - Centralized flow\
- Clear logic\
- Easier debugging

**Cons:** - Tighter coupling to the orchestrator\
- Orchestrator becomes a critical component

------------------------------------------------------------------------

## 🔄 Compensating Transactions

In Sagas, failures are handled using compensating actions.\
Example:\
- Step 1: Reserve inventory\
- Step 2: Charge payment → **Fails**\
- Compensation: Release reserved inventory

Compensation *does not restore the past*, but creates a **new event**
that logically undoes the business action.

------------------------------------------------------------------------

## 🧪 Practical Exercise

Design a Saga for an **E-commerce Checkout** workflow:

### Steps:

1.  CreateOrder\
2.  ReserveInventory\
3.  ProcessPayment\
4.  CreateShipment

### Tasks:

1.  Describe the **happy path** (all steps succeed).\
2.  Describe the **compensation path** if "ProcessPayment" fails.\
3.  Model the saga using:
    -   Choreography (events)\
    -   Orchestration (a central coordinator)\
4.  Define the events or commands involved.

------------------------------------------------------------------------

## ⚠️ When to Use Sagas

### 👍 Use sagas when:

-   You have long-running, multi-step workflows.\
-   Distributed transactions (2PC) are not an option.\
-   You need strong reliability and rollback logic.\
-   Working with microservices or distributed systems.

### 👎 Avoid sagas when:

-   Your workflow is simple.\
-   All operations can be handled inside a single service.\
-   Compensations would be too complex or harmful.

------------------------------------------------------------------------

## 📎 Recommended Resources

-   Chris Richardson --- *Microservices Patterns*\
-   RedHat --- Saga Pattern Guide\
-   Eventuate Tram Saga Framework docs\
-   Netflix Conductor (for orchestration)

------------------------------------------------------------------------

## ✏️ English Writing Feedback

Your request:\
\> "Hi, please provide the instructions for Class 43: Saga Patern.
Provide them as a downloadable Markdown file, and please give feedback
on my English writing."

### Feedback:

-   You wrote **"Saga Patern"**, but the correct spelling is **"Saga
    Pattern"**.\
-   The rest of the sentence is excellent.\
-   A slightly more polished version would be:
    -   "Hi, please provide the instructions for Class 43: Saga Pattern.
        Please provide them as a downloadable Markdown file, and give me
        feedback on my English writing."

Your English is clear, correct, and professional. Great job---just the
small spelling fix!
