# Go NET/HTTP Template

I've created this Go project template which aims to kickstart Go REST projects faster. It includes a structured project layout, custom middleware for logging, PostgreSQL database integration using the `pgx` driver, and support for database migrations using `Flyway`.

## Features

- **Go `net/http`**: Utilizes Go's standard `net/http` package for building HTTP servers.
- **Logging Middleware**: Includes middleware to log incoming requests, responses, and errors for easier debugging and monitoring.
- **Database Service**: Provides integration with a PostgreSQL database using the `pgx` driver, complete with SQL functions and procedures.
- **Database Migrations with Flyway**: Supports schema migrations via Flyway, allowing for easy database versioning.
- **Docker Compose Integration**: Flyway migrations and PostgreSQL services can be run through Docker Compose for a simplified setup.
- **Swagger**: Uses `http-swagger` package for generating docs.

## Prerequisites

- **Go 1.22+**: Make sure you have Go version `1.22` or above installed.
- **PostgreSQL**: A running PostgreSQL instance is required.
- **Docker**: Docker is used for managing database services and migrations with Docker Compose.
- **Flyway**: Used for database schema migrations.

## Getting Started

### 1. Clone the Repository

You can use this template in two ways:

1. **Using GitHub Template**:

   - Click the "Use this template" button on the GitHub repository
   - Create a new repository using this template
   - Clone your new repository:

   ```bash
   git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
   cd YOUR_REPO_NAME
   ```

2. **Direct Clone**:
   ```bash
   git clone https://github.com/sabuhiteymurov/go-http-template.git
   cd go-http-template
   ```

### 2. Install Dependencies

Use `go mod tidy` to install all necessary Go modules.

```bash
go mod tidy
```

### 3. Configure Environment Variables

Create a `.env` file in the root directory based on the provided `.env.example` file.

```bash
cp .env.example .env
```

Update the `.env` file with your PostgreSQL credentials and other configurations.

### 4. Run Migrations with Flyway (via Docker Compose)

Run the following command to set up the database and execute migrations using Docker Compose:

```bash
docker compose up -d
```

This command will start the PostgreSQL container and apply the database migrations using Flyway.

### 5. Run the Server

Once the environment is configured and migrations are complete, you can start the server:

```bash
go run main.go
```

The server will start and be ready to handle requests.

## Project Structure

The project follows a modular structure to keep the code organized and maintainable.

```
go-http-template/
│
├── internal/
│   ├── config/             # Loads environment variables and app configuration
│   ├── repository/         # Database interactions (CRUD operations)
│   ├── handlers/           # Route handlers for different endpoints
│   ├── dto/                # Request/Response DTOs
│   ├── helpers/            # Common helper functions
│   └── middleware/         # Custom middleware (logging, authentication)
│
├── db/
│   ├── migrations/         # SQL migration files (e.g., Flyway)
│   ├── plpgsql/            # PL/pgSQL function or stored procedure files
│   └── seed/               # Optional: Data seeding scripts
│
├── routes/                 # HTTP route definitions
│   └── routes.go
│
├── models/                 # Public models (database models or types)
├── utils/                  # Utility functions and common tools
│
├── docker-compose.yml
├── .env.example
├── go.mod
├── main.go                 # Application entry point
└── README.md
```

## Running Tests

To run tests, use the `go test` command:

```bash
go test ./...
```

This will run all the tests defined across the project.

## Contributing

Contributions are welcomed! Please see the [Contributing Guide](CONTRIBUTING.md) for details on how to submit pull requests, report issues, and contribute to the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, feel free to open an issue or reach out to [me](mailto:hello@sabuhiteymurov.com).

---
