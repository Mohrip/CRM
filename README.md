# CRM Application

This is a simple CRM (Customer Relationship Management) application built using Go. The application provides basic user management functionalities, including creating, retrieving, updating, and deleting user records.

## Project Structure

```
crm-app
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   ├── controllers
│   │   └── user_controller.go  # Handles user-related HTTP requests
│   ├── models
│   │   └── user.go             # Defines the user model
│   ├── routes
│   │   └── routes.go           # Sets up the HTTP routes
│   └── services
│       └── user_service.go     # Contains business logic for user management
├── pkg
│   └── config
│       └── config.go           # Configuration settings and loading
├── go.mod                       # Module definition and dependencies
└── README.md                    # Project documentation
```

## Features

- User Registration
- User Profile Management
- User Deletion

## Getting Started

1. Clone the repository:
   ```
   git clone <repository-url>
   cd crm-app
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the application:
   ```
   go run cmd/main.go
   ```

4. Access the API at `http://localhost:8080`.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.