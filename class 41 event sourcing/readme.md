# Class 41: Event Sourcing Architecture

## 🎯 Objectives

-   Understand the core principles of Event Sourcing.
-   Learn how events become the source of truth.
-   Identify when Event Sourcing is useful and when it is not.
-   Explore how Event Sourcing works together with CQRS.

------------------------------------------------------------------------

## 📘 What Is Event Sourcing?

Event Sourcing is an architectural pattern where **state is not stored
as the current value**, but instead derived from a **sequence of
events**.\
Each change to the application state is stored as an immutable event.

### 🧩 Key Concepts

-   **Event Store:**\
    A database optimized for appending events and retrieving them in
    order.

-   **Events:**\
    Records of something that *already happened* (e.g., `OrderCreated`,
    `OrderPaid`, `OrderShipped`).

-   **Projections:**\
    Read models built from the event stream. They represent the current
    state.

-   **Rehydration:**\
    Rebuilding an object's state by replaying its events.

------------------------------------------------------------------------

## ⚙️ How It Works

1.  A command triggers a business action.
2.  The domain logic generates one or more events.
3.  The events are stored in the event store.
4.  Projections are updated based on these events.
5.  The system state is derived from replaying events or reading
    projections.

------------------------------------------------------------------------

## 🧪 Practical Exercise

Design an Event Sourcing model for an **Order System**:

### Events to define

-   `OrderCreated`
-   `OrderItemAdded`
-   `OrderPaid`
-   `OrderCancelled`
-   `OrderShipped`

### Tasks

1.  Describe what data each event should include.\
2.  Explain how the system reconstructs the order state using replay.\
3.  Define at least one projection (e.g., `OrderSummaryProjection`).\
4.  Explain how Event Sourcing integrates with CQRS in your example.

------------------------------------------------------------------------

## ⚠️ When to Use Event Sourcing

### 👍 Good for:

-   Complex domains with many state transitions.
-   Auditable systems.
-   Systems requiring time-travel debugging.
-   High scalability and distributed consistency.

### 👎 Avoid when:

-   The domain is simple.
-   You don't need historical state.
-   You lack infrastructure for projections and event storage.
-   Developers are unfamiliar with the complexity.

------------------------------------------------------------------------

## 📎 Recommended Resources

-   Greg Young --- *Event Sourcing Basics*
-   EventStoreDB Documentation
-   Martin Fowler --- *Event Sourcing*

------------------------------------------------------------------------

## ✏️ English Writing Feedback

Your request:\
\> "Hi, please provide the instructions for Class 41: Event Sourcing
Architecture. Offer them in a downloadable Markdown file, and please
give feedback on my English writing."

**Feedback:**\
- Your sentence is already clear and natural.\
- "Offer them in a downloadable Markdown file" is correct, but you could
make it slightly more natural by saying:\
- "Provide them as a downloadable Markdown file."\
- Everything else reads smoothly and is grammatically correct.

Great job --- your English is strong and professional!
