# Class 45: API Gateway Pattern

## Objectives
- Understand what an API Gateway is and why it is essential in modern architectures.
- Learn the responsibilities and capabilities of an API Gateway.
- Compare API Gateway vs Load Balancer vs Reverse Proxy.
- Identify common use cases such as routing, throttling, authentication, and protocol translation.
- Build a small API Gateway example (conceptual or with a lightweight tool).
- Recognize anti-patterns and when **not** to use an API Gateway.

---

## 1. Introduction
In microservices architectures, clients would otherwise need to call multiple services directly.  
This leads to:
- Many round trips  
- Complex client logic  
- Security exposure  
- Versioning issues  

The **API Gateway Pattern** provides a single entry point for all clients.  
It acts as a _front door_ to your microservices.

---

## 2. Responsibilities of an API Gateway

### Core Responsibilities
- **Request routing**  
- **Load balancing**  
- **Centralized authentication & authorization**  
- **TLS termination**  
- **Rate limiting / throttling**  
- **Caching**  
- **Monitoring & logging**  
- **Protocol transformation** (REST ↔ gRPC, etc.)

### Extended Responsibilities
- Canary releases  
- A/B testing  
- Request/response shaping  
- Aggregation of multiple services into one endpoint  
- Circuit breaker & retry logic (some gateways support this)

---

## 3. Architecture Diagram (Textual)

```
           +----------------+
Client --> |  API Gateway   | --> Auth Service
           +----------------+
                |    |   |
                |    |   |
                v    v   v
         User Service  Order Service  Payment Service
```

The gateway **receives one request** and then:
- Routes it  
- Validates it  
- Transforms it  
- Aggregates responses if needed  
- Sends back a clean response to the client  

---

## 4. API Gateway vs Other Components

### API Gateway vs Load Balancer
| Feature | API Gateway | Load Balancer |
|--------|-------------|---------------|
| Routing rules | Advanced | Basic |
| Auth | Yes | No |
| Caching | Yes | No |
| Aggregation | Yes | No |
| Protocol conversion | Yes | No |

### API Gateway vs Reverse Proxy
- Both route traffic  
- API Gateway adds authentication, rate limiting, transformations, aggregation, etc.

---

## 5. Example Use Case (Pseudo-code)

### Client call
```
/api/orders/summary
```

### API Gateway logic
```
authenticate()
rate_limit_check()
response_user = call(user_service)
response_orders = call(order_service)
return aggregate(response_user, response_orders)
```

---

## 6. When to Use an API Gateway
✔ Microservices architectures  
✔ When clients should not talk directly to services  
✔ When enforcing security at a single point  
✔ When doing versioning  
✔ For mobile apps needing fewer round trips  

---

## 7. When NOT to Use It
✘ Simple monolith or 2–3 services  
✘ When low latency is critical (gateway adds ~5–10ms)  
✘ If it becomes a single point of failure (avoid by using HA setups)  

---

## 8. Practical Exercise
Create a simple API Gateway with one of these tools:
- **Kong**  
- **Traefik**  
- **NGINX**  
- **KrakenD**  
- **AWS API Gateway**  
- **Go-based custom reverse proxy**

Your gateway must:
- Route 3 different paths to 3 services  
- Add request logging  
- Validate an API key  
- Return a combined response for one endpoint  

Bonus:
- Add rate limiting  
- Add caching  
- Deploy it using Docker  

---

## 9. Additional Resources
- NGINX API Gateway documentation  
- Kong Gateway  
- KrakenD declarative gateway  
- AWS API Gateway official guide  

---

# English Writing Feedback

### 👍 Strengths
- You consistently use polite, concise English.
- Clear structure in your requests.
- Your technical vocabulary is excellent.

### ✍️ Suggestions
- Correct the small typo:  
  - “Api Gateway” → **“API Gateway”**  
- More natural phrasing:  
  - “Hi, please provide the instructions…” is correct.  
  - You can also say:  
    - **“Hi, could you provide the instructions for Class 45: API Gateway and include them as a downloadable Markdown file? Also, please give me feedback on my English writing.”**

Your English is very strong — only minor adjustments needed.
