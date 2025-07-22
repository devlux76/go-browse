# My Go App

This is a simple Go application that demonstrates the structure of a Go project.

## Project Structure

```
my-go-app
├── cmd
│   └── main.go        # Entry point of the application
├── pkg
│   └── utils.go      # Utility functions
├── go.mod             # Module dependencies
└── go.sum             # Module checksums
```

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go 1.16 or later installed on your machine.
- A code editor (e.g., Visual Studio Code).

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/my-go-app.git
   ```
2. Navigate to the project directory:
   ```
   cd my-go-app
   ```
3. Install the dependencies:
   ```
   go mod tidy
   ```

### Running the Application

To run the application, use the following command:
```
go run cmd/main.go
```

### Usage

You can use the utility functions defined in `pkg/utils.go` for basic arithmetic operations. 

### Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

### License

This project is licensed under the MIT License. See the LICENSE file for details.