# Product Management System

A robust and scalable product management system built with Go, featuring user authentication, customer management, and product favorites functionality.

## Features

- **User Management**
  - User registration and authentication
  - Profile management with picture support
  - Account activation/deactivation

- **Product Management**
  - CRUD operations for products
  - Product search with filters
  - Price range filtering
  - Active/Inactive product status

- **Customer Management**
  - Customer profiles with full name parsing
  - Email validation
  - Customer-specific features

- **Favorite Products**
  - Add/Remove favorite products
  - Priority management for favorites
  - Batch operations support
  - Favorite products timeline

## Tech Stack

- **Language**: Go 1.24.2
- **Web Framework**: Gin-Gonic
- **ORM**: GORM
- **Authentication**: JWT (planned)
- **Architecture**: Clean Architecture / DDD

## 📁 Project Structure

```
├── cmd/
│   └── http/              # Application entry points
├── internal/
│   ├── application/       # Application core
│   │   ├── dtos/         # Data Transfer Objects
│   │   ├── entities/     # Domain Entities
│   │   ├── ports/        # Interface Definitions
│   │   ├── shared/       # Shared Components
│   │   └── usecases/     # Business Logic
│   ├── infra/            # Infrastructure Layer
│   │   └── data/         
│   └── presentation/      # Presentation Layer
│       └── http/         
```

## Getting Started

1. **Prerequisites**
   - Go 1.24.2 or higher
   - Git

2. **Installation**
   ```bash
   # Clone the repository
   git clone https://github.com/your-username/products.git

   # Navigate to project directory
   cd products

   # Install dependencies
   go mod download

   # Run the application
   go run cmd/http/main.go
   ```

3. **Environment Variables**
   Create a `.env` file in the root directory:
   ```env
   PORT=8080
   ```

## API Endpoints

### Users
- `POST /users` - Create new user
- `GET /users` - List users
- `GET /users/:id` - Get user details
- `PUT /users/:id/activate` - Activate user
- `PUT /users/:id/deactivate` - Deactivate user
- `PUT /users/:id/profile-picture` - Update profile picture

### Products
- `POST /products` - Create product
- `GET /products` - List products
- `GET /products/:id` - Get product details
- `PUT /products/:id` - Update product
- `DELETE /products/:id` - Delete product

### Customers
- `POST /customers` - Create customer
- `GET /customers` - List customers
- `GET /customers/:id` - Get customer details
- `PUT /customers/:id` - Update customer
- `DELETE /customers/:id` - Delete customer

### Favorite Products
- `POST /customers/:id/favorites` - Add favorite products
- `DELETE /customers/:id/favorites/:productId` - Remove favorite product
- `GET /customers/:id/favorites` - List favorite products
- `PUT /customers/:id/favorites/:productId/priority` - Update favorite priority
