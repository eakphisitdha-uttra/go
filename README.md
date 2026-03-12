# Microservice

A Go-based microservice application with multi-database support and RESTful API endpoints.

## Features

- 🗄️ **Multi-Database Support**: PostgreSQL, MongoDB, MySQL, MSSQL
- 🚀 **Web Framework**: Gin HTTP framework
- 📝 **API Documentation**: Swagger/OpenAPI integration
- 🔐 **JWT Authentication**: JSON Web Token support
- 📊 **Excel Operations**: Excel file processing with excelize
- 🌐 **Web Automation**: Chrome DevTools Protocol support
- 📋 **Structured Logging**: Zap logging library
- 🔧 **Environment Configuration**: dotenv support

## Tech Stack

- **Go 1.23.3**
- **Gin** - HTTP Web Framework
- **Databases**:
  - PostgreSQL (lib/pq)
  - MongoDB (mongo-driver)
  - MySQL (go-sql-driver/mysql)
  - MSSQL (go-mssqldb)
- **Authentication**: golang-jwt/jwt
- **Documentation**: Swagger (swaggo)
- **Logging**: uber-go/zap
- **File Processing**: excelize/v2

## Prerequisites

- Go 1.23.3 or higher
- Docker (for containerized databases)
- Access to all configured databases

## Installation

1. Clone the repository:
```bash
git clone https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
cd YOUR_REPO_NAME
```

2. Install dependencies:
```bash
go mod download
```

3. Copy environment file:
```bash
cp .env.example .env
```

4. Configure your database connections in `.env`

## Configuration

The application uses environment variables for configuration. Key variables include:

- Database connection strings for PostgreSQL, MongoDB, MySQL, MSSQL
- JWT secret key
- Server port (default: 8080)
- Log levels and paths

## Running the Application

### Development

```bash
# Run directly
go run main.go

# Or with air for hot reload
air
```

### Production

```bash
# Build
go build -o microservice main.go

# Run
./microservice
```

### Docker

```bash
# Build image
docker build -t microservice .

# Run container
docker run -p 8080:8080 microservice
```

## API Documentation

Once the server is running, visit:
- Swagger UI: `http://localhost:8080/swagger/index.html`

## Project Structure

```
.
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Dependency checksums
├── .env                 # Environment variables
├── .gitignore           # Git ignore rules
├── databases/           # Database connection modules
│   ├── mongodb/
│   ├── mssql/
│   ├── mysql/
│   └── postgresql/
├── helper/              # Utility functions
├── internals/           # Internal application logic
├── logs/                # Logging configuration
├── responses/           # API response structures
├── routes/              # HTTP route definitions
├── storages/            # File storage (gitignored)
├── templates/           # Template files
└── docker-compose/      # Docker configuration
```

## Database Setup

The application connects to multiple databases. Ensure all databases are running and accessible:

### Using Docker Compose

```bash
docker-compose -f docker-compose/docker-compose.yml up -d
```

## Environment Variables

Create a `.env` file with the following variables:

```env
# PostgreSQL
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database

# MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=your_database

# MySQL
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=your_user
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=your_database

# MSSQL
MSSQL_HOST=localhost
MSSQL_PORT=1433
MSSQL_USER=your_user
MSSQL_PASSWORD=your_password
MSSQL_DATABASE=your_database

# JWT
JWT_SECRET=your_jwt_secret_key

# Server
SERVER_PORT=8080
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions, please open an issue in the GitHub repository.
