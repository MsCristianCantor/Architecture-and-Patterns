# Class 38: Serverless Architecture

## 🎯 Objectives
- Understand the principles of **Serverless Architecture** and its key advantages.
- Learn how Serverless differs from traditional and container-based architectures.
- Explore real-world use cases for serverless applications.
- Implement a simple Serverless example in **Go**.

---

## 🧠 What is Serverless Architecture?
Serverless Architecture is a cloud computing model where developers focus on writing code without managing servers.  
The cloud provider automatically handles **scaling, availability, and infrastructure management**.

Although the name says *“serverless”*, servers still exist — they’re just **abstracted away**.  
You pay only for the actual execution time of your code (e.g., AWS Lambda, Google Cloud Functions, Azure Functions).

### 🔑 Key Characteristics:
- **No server management:** The cloud provider handles provisioning and maintenance.
- **Automatic scaling:** Scales up or down based on demand.
- **Pay-per-execution:** You only pay when your code runs.
- **Event-driven execution:** Functions are triggered by events such as HTTP requests, database updates, or message queue events.

---

## 🧩 UML Diagram
A simplified conceptual view of a Serverless Architecture:

```
+-------------------+          +-------------------+
|   User / Client   |          |   Cloud Provider  |
+-------------------+          +-------------------+
          |                             |
          |  HTTP Request               |
          v                             v
+-------------------+          +---------------------------+
|  API Gateway      |  --->    |  Serverless Function (FaaS)|
+-------------------+          +---------------------------+
                                       |
                                       v
                             +-----------------------+
                             |  Database / Service   |
                             +-----------------------+
```

---

## 💡 Conceptual Example

Imagine a **photo processing system** that runs only when a user uploads a photo:
1. A user uploads a photo to cloud storage.
2. This triggers a Serverless Function.
3. The function resizes the photo and stores it in another bucket.
4. The function terminates automatically after processing.

---

## 💻 Go Example (Simulated Locally)

```go
package main

import (
    "fmt"
)

// Simulated serverless handler
func HandleRequest(event string) string {
    return fmt.Sprintf("Processing event: %s", event)
}

func main() {
    // Simulating invocation by an event
    result := HandleRequest("User uploaded a photo")
    fmt.Println(result)
}
```

In real life, this function would be deployed to a cloud provider (e.g., AWS Lambda) and triggered by an event source.

---

## 🧠 Practical Example
**Task:** Create a Go function that simulates an event-driven flow using Serverless logic.

**Steps:**
1. Define an event structure (e.g., `UserEvent`).
2. Write a function `HandleEvent` that processes the event.
3. Simulate different event types: “UserRegistered”, “OrderCreated”, “FileUploaded”.
4. Print different outputs depending on the event type.

---

## 📚 Additional Resources
- [AWS Lambda Documentation](https://docs.aws.amazon.com/lambda/)
- [Google Cloud Functions Overview](https://cloud.google.com/functions)
- [Azure Functions Guide](https://learn.microsoft.com/en-us/azure/azure-functions/)
- [Serverless Framework](https://www.serverless.com/)

---

## 🚀 Optional Challenge
Deploy a simple Go-based function to **AWS Lambda** or **Google Cloud Functions** using the **Serverless Framework**.  
You can make it respond to an API Gateway event or a file upload.

---

## 🔜 Next Steps
In the next class, you’ll explore **Service-Oriented Architecture (SOA)** and how it differs from microservices and serverless models.
