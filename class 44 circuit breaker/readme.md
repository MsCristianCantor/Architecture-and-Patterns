# Class 43: Circuit Breaker Pattern

## Objectives
- Understand the motivation for the Circuit Breaker pattern.
- Learn how failures propagate in distributed systems.
- Identify the different states of a Circuit Breaker.
- Implement a simple Circuit Breaker in code.
- Recognize when to apply (or avoid) this pattern.

## 1. Introduction
In distributed systems, remote calls fail all the time due to network issues, timeouts, overloaded services, or cascading failures.  
The **Circuit Breaker Pattern** protects your system by preventing repeated calls to a failing dependency and giving it time to recover.

## 2. Problem
Without a Circuit Breaker:
- The application keeps calling a failing service.
- Latency increases.
- Resources get exhausted (threads, DB connections, CPU).
- Failures propagate to other services → cascading failure.

## 3. Solution
A Circuit Breaker monitors remote calls and “opens” when failures pass a threshold.  
This stops calls temporarily, allowing the failing service to recover.

### States
1. **Closed**  
   - Everything works normally.  
   - Calls pass through.  
   - Failures are counted.

2. **Open**  
   - The service is considered unhealthy.  
   - Calls are blocked immediately.  
   - A cooldown timer starts.

3. **Half-Open**  
   - A small number of test requests are allowed.  
   - If they succeed → back to *Closed*.  
   - If they fail → return to *Open*.

```
[ Closed ] → failures → [ Open ] → cooldown → [ Half-Open ] → success → Closed
                                                  ↓
                                               failure
                                                  ↓
                                                Open
```

## 4. Key Concepts
- **Failure Threshold**: How many errors trigger the breaker.
- **Timeout Duration**: How long to stay open.
- **Trial Requests**: Small set of requests in Half-Open.
- **Fallback Strategy**: What to do when the circuit is open.

## 5. Example (pseudo-code)
```pseudo
if state == CLOSED:
    try call_service()
    if fail: count_fail()
    if count_fail > threshold: state = OPEN

if state == OPEN:
    if cooldown_passed: state = HALF_OPEN
    else: return fallback

if state == HALF_OPEN:
    try call_service()
    if success: state = CLOSED; reset_failures()
    else: state = OPEN
```

## 6. When to Use Circuit Breaker
✔ Distributed systems  
✔ Microservices  
✔ External APIs  
✔ Unpredictable networks  

## 7. When NOT to Use It
✘ When failures should NOT interrupt operations (e.g., batch processing).  
✘ When the dependency is reliable and local.  
✘ When retry logic alone is enough.

## 8. Practical Exercise
Create a small service that calls a simulated API with random failures.  
Implement:
- Closed, Open, Half-Open states
- A failure threshold
- Cooldown timer
- A fallback strategy (e.g., cached response)

Repeat using:
- Golang  
- Or any language of your preference

## 9. Additional Resources
- Michael Nygard’s *Release It!* (origin of the pattern)
- Polly (C# library)
- Resilience4j (Java)
- Hystrix (legacy Netflix library)

---

# English Writing Feedback

Your English is already very solid. A few suggestions to improve it even more:

### 👍 Strengths
- Clear and concise sentence structure.
- Your technical writing is very understandable.
- Good use of imperative form (“Provide them”, “Please give feedback”).

### ✍️ Suggested Improvements
- Correct the spelling:  
  - “Saga Patern” → **“Saga Pattern”**  
  - “please provide the instructions” is correct.
- To sound more natural, you can write:  
  - **“Hi, could you provide the instructions for Class 43: Circuit Breaker?”**  
  - **“Please include them as a downloadable Markdown file, and give me feedback on my English writing.”**

Small details — overall your English is excellent.

