# Class 48: Strangler Pattern

## 🎯 Objectives
- Understand the Strangler Pattern and its purpose in system modernization.
- Learn how to migrate legacy systems gradually and safely.
- Identify benefits, trade-offs, and when to apply this pattern.
- Design migration flows using proxies and routing layers.
- Execute a hands-on migration exercise.

---

## 1. Introduction

The **Strangler Pattern** (also known as *Strangler Fig Pattern*) is a modernization technique used to incrementally replace a legacy system with a new one—**without a full rewrite and without downtime**.

It works by:
1. Placing a routing layer in front of the legacy system.
2. Gradually implementing new components.
3. Redirecting traffic to the new components over time.
4. Phasing out the legacy system when no longer needed.

This reduces risk and allows modernization while still serving users.

---

## 2. Origin of the Name

Inspired by the *strangler fig tree*, which grows around its host tree and slowly replaces it.  
Similarly, new functionality “grows” around the legacy system until it can be safely removed.

---

## 3. How the Pattern Works

### **Step 1: Add a Routing Layer**
A proxy, API gateway, or facade becomes the single entry point for all requests.

```
Client → Routing Layer → Legacy
```

### **Step 2: Build New Services**
New modules are developed according to modern standards.

### **Step 3: Redirect Traffic**
Requests for replaced features are routed to new services.

```
Client → Routing Layer → New Service
```

### **Step 4: Decommission Legacy**
Once all features are migrated, the legacy system is removed.

---

## 4. Architecture Diagram

```
            +----------------------+
            |        Client        |
            +----------+-----------+
                       |
                       v
        +-------------------------------+
        |        Routing / Proxy        |
        +-------------------------------+
          /               |             \
   Legacy Module X   New Service Y   New Service Z
```

---

## 5. Benefits

### ✅ Advantages
- Incremental migration
- Minimal risk and downtime
- Feature-by-feature replacement
- Easier testing and rollout
- Smooth coexistence of old and new systems

---

## 6. Drawbacks

### ⚠️ Disadvantages
- More routing complexity
- Dual systems increase operational costs
- Requires clear boundaries and domain understanding
- Migration can take long if not planned properly

---

## 7. When to Use It

Use the Strangler Pattern when:
- Your system is large and hard to rewrite
- You need continuous uptime
- You want gradual, controlled modernization
- You have legacy logic that can be isolated

---

## 8. When *Not* to Use It

Avoid it when:
- The legacy system is too unstable for incremental changes
- Modules can’t be isolated cleanly
- Deadlines require a full replacement
- Your team lacks DevOps / CI/CD maturity

---

## 9. Practical Exercise

### **Exercise 1: Identify a Replacement Candidate**
Choose a module (ex: authentication, catalog, payments).  
Describe how you would extract and modernize it.

---

### **Exercise 2: Traffic Routing Plan**
Explain how you would route:
- 100% of requests initially to legacy  
- Gradual traffic splitting  
- Full cutover to the new service  

---

### **Exercise 3: Clean-Up Strategy**
Define:
- How you will verify that no traffic reaches legacy code  
- How to decommission and archive the old module  

---

## 10. Additional Resources
- Martin Fowler: Strangler Fig Application  
- NGINX: Migrating from Monolith to Microservices  
- AWS / Google Cloud modernization case studies  

---

# 📝 English Writing Feedback

Your English continues to be **clear, structured, and professional**. Excellent work!

### 💪 Strengths
- Strong sentence structure.
- Very clear and direct request.
- Accurate technical vocabulary.

### 🔧 Suggested Improvements
Your original sentence:

> “Hi, please provide the instructions for Class 48: Strangler Pattern. Provide them as a downloadable Markdown file, and please give feedback on my English writing.”

Improved version:

**“Hi, please provide the instructions for Class 48: Strangler Pattern as a downloadable Markdown file, and please include feedback on my English writing.”**

Just a small refinement — your English is already very good.

