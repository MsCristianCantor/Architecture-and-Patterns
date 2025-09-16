# Class 32: Layered Architecture

## Objectives
- Understand the principles of Layered Architecture.
- Learn how to organize applications into layers for better maintainability.
- Implement a simple example of Layered Architecture in Go.
- Apply a practical exercise to consolidate knowledge.

---

## What is Layered Architecture?
Layered Architecture (also known as the **n-tier architecture**) is a design pattern that organizes an application into separate layers, each with distinct responsibilities.  
Common layers include:
- **Presentation layer**: Handles the user interface and user interaction.
- **Business logic layer**: Contains the application rules and processes.
- **Data access layer**: Manages database or external data interaction.

This separation of concerns improves scalability, testability, and maintainability.

---

## UML Diagram

```
+----------------------+
|   Presentation       |
+----------------------+
          |
          v
+----------------------+
|   Business Logic     |
+----------------------+
          |
          v
+----------------------+
|   Data Access        |
+----------------------+
```

---

## Conceptual Example
Imagine an online store:
- **Presentation layer**: A web page or API that receives a request to buy a product.
- **Business logic layer**: Validates if the product is in stock and applies discounts.
- **Data access layer**: Reads product information from the database and updates the stock.

---

## Go Example

```go
package main

import "fmt"

// Data Access Layer
type ProductRepository struct{}

func (r *ProductRepository) GetProduct(productID int) string {
	return "Laptop"
}

// Business Logic Layer
type ProductService struct {
	repo *ProductRepository
}

func (s *ProductService) ProcessOrder(productID int) string {
	product := s.repo.GetProduct(productID)
	return fmt.Sprintf("Order processed for product: %s", product)
}

// Presentation Layer
type ProductController struct {
	service *ProductService
}

func (c *ProductController) HandleOrder(productID int) {
	result := c.service.ProcessOrder(productID)
	fmt.Println(result)
}

func main() {
	repo := &ProductRepository{}
	service := &ProductService{repo: repo}
	controller := &ProductController{service: service}

	controller.HandleOrder(1)
}
```

---

## Practical Example
- Build a small application for **user registration**:
  - **Presentation layer**: Console input/output for the user.
  - **Business layer**: Validate email format and password strength.
  - **Data access layer**: Store user info in memory or a simple file.

---

## Additional Resources
- [Layered Architecture Explained](https://martinfowler.com/architecture/)
- *Patterns of Enterprise Application Architecture* by Martin Fowler
- Go blog examples of modular applications

---

## Optional Challenge
- Extend the practical example by adding a **logging layer** to track application operations.
- Implement error handling across the layers.
- Replace the in-memory data storage with a real database connection.

---

## Next Steps
Next, we will explore **Clean Architecture**, which builds on layered principles but provides stricter rules and independence from frameworks.
