# Go Microservice

[![Go](https://img.shields.io/badge/Go-1.23.3-blue.svg)](https://golang.org/)
[![Gin](https://img.shields.io/badge/Gin-1.10.0-blue.svg)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-latest-blue.svg)](https://www.postgresql.org/)
[![MongoDB](https://img.shields.io/badge/MongoDB-latest-blue.svg)](https://www.mongodb.com/)
[![MySQL](https://img.shields.io/badge/MySQL-latest-blue.svg)](https://www.mysql.com/)
[![MSSQL](https://img.shields.io/badge/MSSQL-latest-blue.svg)](https://www.microsoft.com/sql-server)

## 📋 สารบัญ

- [🚀 เริ่มต้นใช้งาน](#-เริ่มต้นใช้งาน)
- [📐 สถาปัตยกรรม](#-สถาปัตยกรรม)
- [📂 โครงสร้างโปรเจค](#-โครงสร้างโปรเจค)
- [🧩 แนวทางการเพิ่ม Feature ใหม่](#-แนวทางการเพิ่ม-feature-ใหม่)
- [🛠️ แนวทางการพัฒนา](#-แนวทางการพัฒนา)
- [🔒 ความปลอดภัย](#-ความปลอดภัย)
- [⚡ การปรับแต่งประสิทธิภาพ](#-การปรับแต่งประสิทธิภาพ)
- [📚 เทคโนโลยีหลัก](#-เทคโนโลยีหลัก)
- [📝 แนวทางการทดสอบ](#-แนวทางการทดสอบ)
- [🔄 CI/CD](#-cicd)
- [📘 เอกสารอ้างอิงเพิ่มเติม](#-เอกสารอ้างอิงเพิ่มเติม)

## 🚀 เริ่มต้นใช้งาน

### ความต้องการของระบบ

- Go 1.23.3 หรือสูงกว่า
- Docker และ Docker Compose (สำหรับฐานข้อมูลในคอนเทนเนอร์)
- ความเข้าใจพื้นฐานเกี่ยวกับ Go, RESTful APIs และฐานข้อมูล

### การติดตั้ง

```bash
# โคลน repository
git clone https://github.com/eakphisitdha-uttra/go.git
cd go

# ติดตั้ง dependencies
go mod download

# สร้างไฟล์ .env (ดูรายละเอียดในหัวข้อ "Environment Variables")
cp .env.example .env
```

### การรันในโหมดพัฒนา

```bash
# รันโดยตรง
go run main.go

# หรือใช้ air สำหรับ hot reload (ต้องติดตั้ง air ก่อน)
air
```

เซิร์ฟเวอร์จะทำงานที่ http://localhost:8080

### การสร้างและใช้งานในโหมด Production

```bash
# สร้าง binary สำหรับ production
go build -o microservice main.go

# รัน production server
./microservice
```

### การใช้งานกับ Docker

```bash
# Build image
docker build -t microservice .

# Run container
docker run -p 8080:8080 microservice

# หรือใช้ docker-compose
docker-compose -f docker-compose/docker-compose.yml up -d
```

### Environment Variables

สร้างไฟล์ `.env` ที่ root ของโปรเจคและกำหนดค่าตามนี้:

```env
# PostgreSQL Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=your_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=your_database
POSTGRES_SSLMODE=disable

# MongoDB Configuration
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=your_database
MONGODB_TIMEOUT=10s

# MySQL Configuration
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=your_user
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=your_database
MYSQL_CHARSET=utf8mb4

# MSSQL Configuration
MSSQL_HOST=localhost
MSSQL_PORT=1433
MSSQL_USER=your_user
MSSQL_PASSWORD=your_password
MSSQL_DATABASE=your_database
MSSQL_ENCRYPT=disable

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRES_IN=24h

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
GIN_MODE=release

# Logging Configuration
LOG_LEVEL=info
LOG_FILE_PATH=logs/
LOG_MAX_SIZE=100
LOG_MAX_BACKUPS=3
LOG_MAX_AGE=28

# Swagger Configuration
SWAGGER_HOST=localhost:8080
SWAGGER_BASE_PATH=/api/v1
```

## 📐 สถาปัตยกรรม

โปรเจคนี้ใช้แนวทาง **Layered Architecture with Database Abstraction** โดยมีลักษณะดังนี้:

### Layered Architecture
โครงสร้างแบบชั้นๆ ที่แยกความรับผิดชอบอย่างชัดเจน:

- **Presentation Layer**: Routes, Controllers, Middleware
- **Business Logic Layer**: Services, Domain Logic
- **Data Access Layer**: Database repositories, Models
- **Infrastructure Layer**: Database connections, External services

### Database Abstraction Pattern
ใช้ pattern เดียวกันสำหรับทุกฐานข้อมูลเพื่อให้ง่ายต่อการสลับและบำรุงรักษา:

```go
// Interface ที่ใช้ร่วมกัน
type Database interface {
    Connect() error
    Close() error
    GetConnection() interface{}
}

// Implementation สำหรับแต่ละฐานข้อมูล
type PostgreSQL struct {
    connection *sql.DB
}
type MongoDB struct {
    client *mongo.Client
}
```

### Dependency Injection
ใช้ pattern ของการ inject dependencies ผ่าน constructor เพื่อให้ง่ายต่อการทดสอบและบำรุงรักษา

### การจัดการการเชื่อมต่อฐานข้อมูล
- **Connection Pooling**: ใช้ connection pool สำหรับ SQL databases
- **Context-aware Operations**: ใช้ context สำหรับ timeout และ cancellation
- **Graceful Shutdown**: ปิดการเชื่อมต่ออย่างเรียบร้อยเมื่อหยุดเซิร์ฟเวอร์

## 📂 โครงสร้างโปรเจค

```
.
├── main.go                    # Entry point ของแอปพลิเคชัน
├── go.mod                     # Go module definition
├── go.sum                     # Dependency checksums
├── .env                       # Environment variables (local)
├── .env.example               # Environment variables template
├── .gitignore                 # Git ignore rules
├── Dockerfile                 # Docker image definition
├── docker-compose/            # Docker Compose configurations
│   ├── docker-compose.yml     # Main compose file
│   ├── postgres.yml           # PostgreSQL service
│   ├── mongodb.yml            # MongoDB service
│   ├── mysql.yml              # MySQL service
│   └── mssql.yml              # MSSQL service
├── databases/                 # Database connection modules
│   ├── mongodb/               # MongoDB connection
│   │   ├── connection.go      # MongoDB connection logic
│   │   └── models.go          # MongoDB models
│   ├── mssql/                 # MSSQL connection
│   │   ├── connection.go      # MSSQL connection logic
│   │   └── models.go          # MSSQL models
│   ├── mysql/                 # MySQL connection
│   │   ├── connection.go      # MySQL connection logic
│   │   └── models.go          # MySQL models
│   └── postgresql/            # PostgreSQL connection
│       ├── connection.go      # PostgreSQL connection logic
│       └── models.go          # PostgreSQL models
├── helper/                    # Utility functions
│   ├── auth.go                # Authentication helpers
│   ├── validation.go          # Input validation
│   ├── response.go            # Response formatting
│   └── utils.go               # General utilities
├── internals/                 # Internal application logic
│   ├── controllers/           # HTTP controllers
│   ├── middleware/            # HTTP middleware
│   ├── models/                # Domain models
│   ├── services/              # Business logic services
│   └── repositories/          # Data access repositories
├── logs/                      # Logging configuration
│   ├── logger.go              # Logger setup
│   ├── error.log              # Error logs (gitignored)
│   ├── info.log               # Info logs (gitignored)
│   └── debug.log              # Debug logs (gitignored)
├── responses/                 # API response structures
│   ├── api_response.go        # Standard API response format
│   ├── error_response.go      # Error response format
│   └── success_response.go    # Success response format
├── routes/                    # HTTP route definitions
│   ├── routes.go              # Route setup
│   ├── api/                   # API routes
│   │   ├── v1/                # API version 1
│   │   │   ├── auth.go        # Authentication routes
│   │   │   ├── users.go       # User management routes
│   │   │   └── health.go      # Health check routes
│   │   └── middleware/        # Route middleware
│   └── swagger/               # Swagger documentation
├── storages/                  # File storage (gitignored)
├── templates/                 # Template files
└── docs/                      # Documentation
    ├── api.md                 # API documentation
    └── deployment.md          # Deployment guide
```

## 🧩 แนวทางการเพิ่ม Feature ใหม่

### 1. สร้างโครงสร้างฟีเจอร์

```bash
# สร้าง directories สำหรับ feature ใหม่
mkdir -p internals/models/[feature-name]
mkdir -p internals/services/[feature-name]
mkdir -p internals/repositories/[feature-name]
mkdir -p internals/controllers/[feature-name]
mkdir -p routes/api/v1/[feature-name]
```

### 2. กำหนด Models (สร้างไฟล์ `internals/models/[feature-name]/model.go`)

```go
package [feature-name]

import (
    "time"
    "github.com/google/uuid"
)

// YourEntity represents the [feature-name] entity
type YourEntity struct {
    ID        string    `json:"id" db:"id"`
    Name      string    `json:"name" db:"name"`
    Email     string    `json:"email" db:"email"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateEntityRequest represents request for creating entity
type CreateEntityRequest struct {
    Name  string `json:"name" validate:"required,min=3,max=100"`
    Email string `json:"email" validate:"required,email"`
}

// UpdateEntityRequest represents request for updating entity
type UpdateEntityRequest struct {
    Name  string `json:"name" validate:"omitempty,min=3,max=100"`
    Email string `json:"email" validate:"omitempty,email"`
}

// EntityResponse represents response format
type EntityResponse struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the table name for database
func (YourEntity) TableName() string {
    return "[table_name]"
}
```

### 3. สร้าง Repository (สร้างไฟล์ `internals/repositories/[feature-name]/repository.go`)

```go
package [feature-name]

import (
    "context"
    "database/sql"
    "errors"
    "fmt"
    "github.com/jmoiron/sqlx"
    "github.com/google/uuid"
    "microservice/internals/models/[feature-name]"
)

// Repository interface defines database operations
type Repository interface {
    Create(ctx context.Context, entity *models.YourEntity) error
    GetByID(ctx context.Context, id string) (*models.YourEntity, error)
    GetAll(ctx context.Context, limit, offset int) ([]*models.YourEntity, error)
    Update(ctx context.Context, id string, entity *models.YourEntity) error
    Delete(ctx context.Context, id string) error
}

// repository implements Repository interface
type repository struct {
    db *sqlx.DB
}

// NewRepository creates a new repository instance
func NewRepository(db *sqlx.DB) Repository {
    return &repository{db: db}
}

// Create creates a new entity
func (r *repository) Create(ctx context.Context, entity *models.YourEntity) error {
    query := `
        INSERT INTO [table_name] (id, name, email, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
    `
    
    _, err := r.db.ExecContext(ctx, query, 
        entity.ID, entity.Name, entity.Email, entity.CreatedAt, entity.UpdatedAt)
    
    return err
}

// GetByID retrieves an entity by ID
func (r *repository) GetByID(ctx context.Context, id string) (*models.YourEntity, error) {
    query := `
        SELECT id, name, email, created_at, updated_at
        FROM [table_name]
        WHERE id = $1
    `
    
    var entity models.YourEntity
    err := r.db.GetContext(ctx, &entity, query, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("entity not found")
        }
        return nil, err
    }
    
    return &entity, nil
}

// GetAll retrieves all entities with pagination
func (r *repository) GetAll(ctx context.Context, limit, offset int) ([]*models.YourEntity, error) {
    query := `
        SELECT id, name, email, created_at, updated_at
        FROM [table_name]
        ORDER BY created_at DESC
        LIMIT $1 OFFSET $2
    `
    
    var entities []*models.YourEntity
    err := r.db.SelectContext(ctx, &entities, query, limit, offset)
    return entities, err
}

// Update updates an existing entity
func (r *repository) Update(ctx context.Context, id string, entity *models.YourEntity) error {
    query := `
        UPDATE [table_name]
        SET name = $1, email = $2, updated_at = $3
        WHERE id = $4
    `
    
    result, err := r.db.ExecContext(ctx, query, 
        entity.Name, entity.Email, entity.UpdatedAt, id)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("entity not found")
    }
    
    return nil
}

// Delete deletes an entity by ID
func (r *repository) Delete(ctx context.Context, id string) error {
    query := `DELETE FROM [table_name] WHERE id = $1`
    
    result, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return fmt.Errorf("entity not found")
    }
    
    return nil
}
```

### 4. สร้าง Service (สร้างไฟล์ `internals/services/[feature-name]/service.go`)

```go
package [feature-name]

import (
    "context"
    "errors"
    "time"
    "github.com/google/uuid"
    "microservice/internals/models/[feature-name]"
    "microservice/internals/repositories/[feature-name]"
)

// Service interface defines business logic operations
type Service interface {
    Create(ctx context.Context, req *models.CreateEntityRequest) (*models.EntityResponse, error)
    GetByID(ctx context.Context, id string) (*models.EntityResponse, error)
    GetAll(ctx context.Context, limit, offset int) ([]*models.EntityResponse, error)
    Update(ctx context.Context, id string, req *models.UpdateEntityRequest) (*models.EntityResponse, error)
    Delete(ctx context.Context, id string) error
}

// service implements Service interface
type service struct {
    repo repositories.Repository
}

// NewService creates a new service instance
func NewService(repo repositories.Repository) Service {
    return &service{repo: repo}
}

// Create creates a new entity
func (s *service) Create(ctx context.Context, req *models.CreateEntityRequest) (*models.EntityResponse, error) {
    // Business logic validation
    if req.Name == "" {
        return nil, errors.New("name is required")
    }
    
    // Create entity
    now := time.Now()
    entity := &models.YourEntity{
        ID:        uuid.New().String(),
        Name:      req.Name,
        Email:     req.Email,
        CreatedAt: now,
        UpdatedAt: now,
    }
    
    // Save to database
    err := s.repo.Create(ctx, entity)
    if err != nil {
        return nil, err
    }
    
    // Return response
    return &models.EntityResponse{
        ID:        entity.ID,
        Name:      entity.Name,
        Email:     entity.Email,
        CreatedAt: entity.CreatedAt,
        UpdatedAt: entity.UpdatedAt,
    }, nil
}

// GetByID retrieves an entity by ID
func (s *service) GetByID(ctx context.Context, id string) (*models.EntityResponse, error) {
    entity, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    return &models.EntityResponse{
        ID:        entity.ID,
        Name:      entity.Name,
        Email:     entity.Email,
        CreatedAt: entity.CreatedAt,
        UpdatedAt: entity.UpdatedAt,
    }, nil
}

// GetAll retrieves all entities
func (s *service) GetAll(ctx context.Context, limit, offset int) ([]*models.EntityResponse, error) {
    entities, err := s.repo.GetAll(ctx, limit, offset)
    if err != nil {
        return nil, err
    }
    
    responses := make([]*models.EntityResponse, len(entities))
    for i, entity := range entities {
        responses[i] = &models.EntityResponse{
            ID:        entity.ID,
            Name:      entity.Name,
            Email:     entity.Email,
            CreatedAt: entity.CreatedAt,
            UpdatedAt: entity.UpdatedAt,
        }
    }
    
    return responses, nil
}

// Update updates an existing entity
func (s *service) Update(ctx context.Context, id string, req *models.UpdateEntityRequest) (*models.EntityResponse, error) {
    // Get existing entity
    entity, err := s.repo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Update fields
    if req.Name != "" {
        entity.Name = req.Name
    }
    if req.Email != "" {
        entity.Email = req.Email
    }
    entity.UpdatedAt = time.Now()
    
    // Save to database
    err = s.repo.Update(ctx, id, entity)
    if err != nil {
        return nil, err
    }
    
    return &models.EntityResponse{
        ID:        entity.ID,
        Name:      entity.Name,
        Email:     entity.Email,
        CreatedAt: entity.CreatedAt,
        UpdatedAt: entity.UpdatedAt,
    }, nil
}

// Delete deletes an entity
func (s *service) Delete(ctx context.Context, id string) error {
    return s.repo.Delete(ctx, id)
}
```

### 5. สร้าง Controller (สร้างไฟล์ `internals/controllers/[feature-name]/controller.go`)

```go
package [feature-name]

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
    "microservice/helper"
    "microservice/internals/models/[feature-name]"
    "microservice/internals/services/[feature-name]"
    "microservice/responses"
)

// Controller handles HTTP requests
type Controller struct {
    service services.Service
}

// NewController creates a new controller instance
func NewController(service services.Service) *Controller {
    return &Controller{service: service}
}

// Create handles POST /api/v1/[feature-name]
func (c *Controller) Create(ctx *gin.Context) {
    var req models.CreateEntityRequest
    
    // Bind request
    if err := ctx.ShouldBindJSON(&req); err != nil {
        responses.Error(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
        return
    }
    
    // Validate request
    if err := helper.ValidateStruct(&req); err != nil {
        responses.Error(ctx, http.StatusBadRequest, "Validation failed", err.Error())
        return
    }
    
    // Create entity
    entity, err := c.service.Create(ctx.Request.Context(), &req)
    if err != nil {
        responses.Error(ctx, http.StatusInternalServerError, "Failed to create entity", err.Error())
        return
    }
    
    responses.Success(ctx, http.StatusCreated, "Entity created successfully", entity)
}

// GetByID handles GET /api/v1/[feature-name]/:id
func (c *Controller) GetByID(ctx *gin.Context) {
    id := ctx.Param("id")
    
    // Get entity
    entity, err := c.service.GetByID(ctx.Request.Context(), id)
    if err != nil {
        responses.Error(ctx, http.StatusNotFound, "Entity not found", err.Error())
        return
    }
    
    responses.Success(ctx, http.StatusOK, "Entity retrieved successfully", entity)
}

// GetAll handles GET /api/v1/[feature-name]
func (c *Controller) GetAll(ctx *gin.Context) {
    // Parse pagination parameters
    limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
    if err != nil || limit <= 0 {
        limit = 10
    }
    
    offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
    if err != nil || offset < 0 {
        offset = 0
    }
    
    // Get entities
    entities, err := c.service.GetAll(ctx.Request.Context(), limit, offset)
    if err != nil {
        responses.Error(ctx, http.StatusInternalServerError, "Failed to retrieve entities", err.Error())
        return
    }
    
    responses.Success(ctx, http.StatusOK, "Entities retrieved successfully", entities)
}

// Update handles PUT /api/v1/[feature-name]/:id
func (c *Controller) Update(ctx *gin.Context) {
    id := ctx.Param("id")
    var req models.UpdateEntityRequest
    
    // Bind request
    if err := ctx.ShouldBindJSON(&req); err != nil {
        responses.Error(ctx, http.StatusBadRequest, "Invalid request body", err.Error())
        return
    }
    
    // Validate request
    if err := helper.ValidateStruct(&req); err != nil {
        responses.Error(ctx, http.StatusBadRequest, "Validation failed", err.Error())
        return
    }
    
    // Update entity
    entity, err := c.service.Update(ctx.Request.Context(), id, &req)
    if err != nil {
        responses.Error(ctx, http.StatusInternalServerError, "Failed to update entity", err.Error())
        return
    }
    
    responses.Success(ctx, http.StatusOK, "Entity updated successfully", entity)
}

// Delete handles DELETE /api/v1/[feature-name]/:id
func (c *Controller) Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    
    // Delete entity
    err := c.service.Delete(ctx.Request.Context(), id)
    if err != nil {
        responses.Error(ctx, http.StatusInternalServerError, "Failed to delete entity", err.Error())
        return
    }
    
    responses.Success(ctx, http.StatusOK, "Entity deleted successfully", nil)
}
```

### 6. สร้าง Routes (สร้างไฟล์ `routes/api/v1/[feature-name]/routes.go`)

```go
package [feature-name]

import (
    "github.com/gin-gonic/gin"
    "microservice/internals/controllers/[feature-name]"
    "microservice/internals/middleware"
    "microservice/internals/repositories/[feature-name]"
    "microservice/internals/services/[feature-name]"
    "github.com/jmoiron/sqlx"
)

// RegisterRoutes registers [feature-name] routes
func RegisterRoutes(router *gin.RouterGroup, db *sqlx.DB) {
    // Create dependencies
    repo := repositories.NewRepository(db)
    service := services.NewService(repo)
    controller := controllers.NewController(service)
    
    // Create route group
    featureRoutes := router.Group("/[feature-name]")
    {
        // Apply authentication middleware
        featureRoutes.Use(middleware.AuthMiddleware())
        
        // CRUD routes
        featureRoutes.POST("", controller.Create)
        featureRoutes.GET("", controller.GetAll)
        featureRoutes.GET("/:id", controller.GetByID)
        featureRoutes.PUT("/:id", controller.Update)
        featureRoutes.DELETE("/:id", controller.Delete)
    }
}
```

### 7. อัปเดต Main Routes (แก้ไข `routes/routes.go`)

```go
package routes

import (
    "database/sql"
    
    "github.com/gin-gonic/gin"
    "github.com/jmoiron/sqlx"
    "microservice/databases/postgresql"
    "microservice/routes/api/v1/[feature-name]"
    "microservice/routes/api/v1/auth"
    "microservice/routes/api/v1/health"
)

// SetupRouter configures and returns the Gin router
func SetupRouter(pg *sql.DB, mg interface{}) *gin.Engine {
    router := gin.Default()
    
    // Convert to sqlx.DB
    pgx := sqlx.NewDb(pg, "postgres")
    
    // API v1 routes
    v1 := router.Group("/api/v1")
    {
        // Health check routes
        health.RegisterRoutes(v1)
        
        // Authentication routes
        auth.RegisterRoutes(v1, pgx)
        
        // [Feature-name] routes
        [feature-name].RegisterRoutes(v1, pgx)
    }
    
    // Swagger documentation
    router.Static("/swagger", "./docs/swagger")
    
    return router
}
```

## 🛠️ แนวทางการพัฒนา

### การใช้งาน Gin Framework

#### Middleware Development
```go
// Custom middleware example
func LoggerMiddleware() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
            param.ClientIP,
            param.TimeStamp.Format(time.RFC1123),
            param.Method,
            param.Path,
            param.Request.Proto,
            param.StatusCode,
            param.Latency,
            param.Request.UserAgent(),
            param.ErrorMessage,
        )
    })
}

// Authentication middleware
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
            c.Abort()
            return
        }
        
        // Validate token logic here
        // ...
        
        c.Next()
    }
}
```

#### Error Handling
```go
// Centralized error handling
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        
        // Check for errors
        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err
            
            // Log error
            log.Printf("Error: %v", err)
            
            // Return error response
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Internal server error",
                "message": err.Error(),
            })
        }
    }
}
```

### การจัดการฐานข้อมูล

#### Connection Pooling
```go
// PostgreSQL connection pool configuration
func Connect() (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("POSTGRES_HOST"),
        os.Getenv("POSTGRES_PORT"),
        os.Getenv("POSTGRES_USER"),
        os.Getenv("POSTGRES_PASSWORD"),
        os.Getenv("POSTGRES_DB"),
        os.Getenv("POSTGRES_SSLMODE"),
    )
    
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }
    
    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)
    
    // Test connection
    if err := db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}
```

#### Transactions
```go
// Transaction handling
func (r *repository) CreateWithTransaction(ctx context.Context, entities []*models.YourEntity) error {
    tx, err := r.db.BeginTxx(ctx, nil)
    if err != nil {
        return err
    }
    
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()
    
    for _, entity := range entities {
        if err := r.createInTx(ctx, tx, entity); err != nil {
            return err
        }
    }
    
    return nil
}
```

### การทำงานกับ Context

#### Context-aware Operations
```go
// Use context for timeout and cancellation
func (s *service) Create(ctx context.Context, req *models.CreateEntityRequest) (*models.EntityResponse, error) {
    // Create context with timeout
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    // Check if context is cancelled
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        // Continue with operation
        return s.createEntity(ctx, req)
    }
}
```

### การจัดการ Logging

#### Structured Logging with Zap
```go
// Logger configuration
func InitLogger() *zap.Logger {
    config := zap.NewProductionConfig()
    config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
    config.OutputPaths = []string{"stdout", "logs/app.log"}
    config.ErrorOutputPaths = []string{"stderr", "logs/error.log"}
    
    logger, err := config.Build()
    if err != nil {
        log.Fatal("Failed to initialize logger:", err)
    }
    
    return logger
}

// Usage in services
func (s *service) Create(ctx context.Context, req *models.CreateEntityRequest) (*models.EntityResponse, error) {
    logger.Info("Creating entity",
        zap.String("name", req.Name),
        zap.String("email", req.Email),
        zap.Time("timestamp", time.Now()),
    )
    
    // Business logic...
    
    logger.Info("Entity created successfully",
        zap.String("id", entity.ID),
        zap.Duration("duration", time.Since(start)),
    )
    
    return entity, nil
}
```

## 🔒 ความปลอดภัย

### การป้องกัน SQL Injection
- ใช้ prepared statements และ parameterized queries เสมอ
- หลีกเลี่ยงการสร้าง SQL queries ด้วยการต่อ string
- ใช้ ORM หรือ query builder ที่มีการป้องกันในตัว

```go
// ✅ ใช้ prepared statements
query := "SELECT * FROM users WHERE id = $1"
err := db.GetContext(ctx, &user, query, userID)

// ❌ หลีกเลี่ยงการต่อ string
query := fmt.Sprintf("SELECT * FROM users WHERE id = '%s'", userID) // อันตราย!
```

### การจัดการ JWT Authentication
- ใช้ signing keys ที่แข็งแกร่ง (256-bit ขึ้นไป)
- กำหนดเวลาหมดอายุที่เหมาะสม
- เก็บ JWT secrets ใน environment variables หรือ secret management systems

```go
// JWT token generation
func GenerateToken(userID string, secretKey string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
        "iat":     time.Now().Unix(),
    })
    
    return token.SignedString([]byte(secretKey))
}
```

### การป้องกัน CORS
- กำหนด CORS policy ที่ชัดเจน
- จำกัด origins ที่อนุญาตให้เข้าถึง
- ใช้ HTTPS ใน production

```go
// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", os.Getenv("ALLOWED_ORIGINS"))
        c.Header("Access-Control-Allow-Credentials", "true")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}
```

### การป้องกัน Rate Limiting
- ใช้ rate limiting middleware เพื่อป้องกัน DDoS attacks
- กำหนดขีดจำกัดตาม IP address หรือ user ID

```go
// Rate limiting middleware
func RateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(rate.Limit(100), 200) // 100 requests per second, burst of 200
    
    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "error": "Rate limit exceeded",
            })
            c.Abort()
            return
        }
        
        c.Next()
    }
}
```

### การป้องกัน Input Validation
- ตรวจสอบและ sanitize ข้อมูลที่นำเข้าทุกครั้ง
- ใช้ validation libraries เช่น go-playground/validator
- จำกัดขนาดของ request body

```go
// Input validation
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required,min=3,max=100"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8,max=128"`
}

func (c *Controller) Create(ctx *gin.Context) {
    var req CreateUserRequest
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
        return
    }
    
    if err := validator.New().Struct(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": err})
        return
    }
    
    // Continue with business logic...
}
```

## ⚡ การปรับแต่งประสิทธิภาพ

### การใช้ Connection Pooling
- กำหนดขนาด connection pool ที่เหมาะสม
- ใช้ connection pooling สำหรับทุกฐานข้อมูล
- ตรวจสอบและปรับแต่งค่า pool settings ตาม workload

### การใช้ Caching
- ใช้ Redis หรือใน-memory cache สำหรับข้อมูลที่เข้าถึงบ่อย
- กำหนด cache expiration ที่เหมาะสม
- ใช้ cache invalidation strategy ที่เหมาะสม

```go
// Cache service example
type CacheService interface {
    Get(key string) (interface{}, error)
    Set(key string, value interface{}, expiration time.Duration) error
    Delete(key string) error
}

// Repository with caching
func (r *repository) GetByIDWithCache(ctx context.Context, id string) (*models.YourEntity, error) {
    cacheKey := fmt.Sprintf("entity:%s", id)
    
    // Try to get from cache first
    if cached, err := r.cache.Get(cacheKey); err == nil {
        return cached.(*models.YourEntity), nil
    }
    
    // Get from database
    entity, err := r.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }
    
    // Cache the result
    r.cache.Set(cacheKey, entity, 5*time.Minute)
    
    return entity, nil
}
```

### การใช้ Background Workers
- ใช้ goroutines สำหรับงานที่ต้องการการประมวลผลนาน
- ใช้ worker pools สำหรับจัดการงานจำนวนมาก
- ใช้ channels สำหรับการสื่อสารระหว่าง goroutines

```go
// Worker pool implementation
type WorkerPool struct {
    workers   int
    jobQueue  chan Job
    workerPool chan chan Job
    quit      chan bool
}

type Job struct {
    ID       string
    Function func() error
}

func NewWorkerPool(workers int) *WorkerPool {
    return &WorkerPool{
        workers:   workers,
        jobQueue:  make(chan Job, 100),
        workerPool: make(chan chan Job, workers),
        quit:      make(chan bool),
    }
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.workers; i++ {
        worker := NewWorker(wp.workerPool)
        worker.Start()
    }
    
    go wp.dispatch()
}

func (wp *WorkerPool) AddJob(job Job) {
    select {
    case wp.jobQueue <- job:
    default:
        log.Println("Job queue is full")
    }
}
```

### การใช้ Database Indexing
- สร้าง indexes สำหรับ columns ที่ใช้ใน WHERE clauses
- ใช้ composite indexes สำหรับ queries ที่ซับซ้อน
- ตรวจสอบ query performance ด้วย EXPLAIN ANALYZE

### การใช้ Pagination
- ใช้ limit/offset หรือ cursor-based pagination
- จำกัดขนาดของ response data
- ใช้ compression สำหรับ large responses

```go
// Cursor-based pagination
type PaginationRequest struct {
    Limit  int    `json:"limit"`
    Cursor string `json:"cursor"`
}

type PaginationResponse struct {
    Data       interface{} `json:"data"`
    NextCursor string      `json:"next_cursor,omitempty"`
    HasMore    bool        `json:"has_more"`
}

func (r *repository) GetAllWithPagination(ctx context.Context, req PaginationRequest) (*PaginationResponse, error) {
    query := "SELECT * FROM entities WHERE created_at < $1 ORDER BY created_at DESC LIMIT $2"
    
    var entities []*models.YourEntity
    err := r.db.SelectContext(ctx, &entities, query, req.Cursor, req.Limit+1)
    if err != nil {
        return nil, err
    }
    
    response := &PaginationResponse{
        Data: entities,
    }
    
    if len(entities) > req.Limit {
        response.HasMore = true
        response.Data = entities[:req.Limit]
        response.NextCursor = entities[req.Limit-1].CreatedAt.Format(time.RFC3339)
    }
    
    return response, nil
}
```

## 📚 เทคโนโลยีหลัก

### Web Framework
- **Gin 1.10.0**: HTTP web framework ที่มีประสิทธิภาพสูงสำหรับ Go
- **Gin Router**: HTTP router ที่รวดเร็วและมีความยืดหยุ่น

### Database Drivers
- **lib/pq**: PostgreSQL driver สำหรับ Go
- **go-mssqldb**: Microsoft SQL Server driver
- **go-sql-driver/mysql**: MySQL driver
- **mongo-driver**: MongoDB official driver

### Authentication & Security
- **golang-jwt/jwt**: JWT implementation สำหรับ Go
- **bcrypt**: Password hashing

### HTTP Client
- **Axios-equivalent**: ใช้ net/http และ http client ของ Go

### Logging
- **uber-go/zap**: Structured logging library ที่มีประสิทธิภาพสูง

### Documentation
- **swaggo/gin-swagger**: Swagger/OpenAPI documentation สำหรับ Gin
- **swaggo/files**: Static file serving สำหรับ Swagger UI

### File Processing
- **xuri/excelize/v2**: Excel file processing library

### Validation
- **go-playground/validator**: Struct validation library

### Testing
- **stretchr/testify**: Testing framework สำหรับ Go
- **gomock**: Mocking framework

### Utilities
- **google/uuid**: UUID generation
- **joho/godotenv**: Environment variable loading

## 📝 แนวทางการทดสอบ

### Unit Testing
- ใช้ testify และ Go's built-in testing package
- ทดสอบ business logic ใน services layer
- ทดสอบ utility functions และ helpers

```go
// Service unit test example
func TestService_Create(t *testing.T) {
    // Mock repository
    mockRepo := &MockRepository{}
    service := NewService(mockRepo)
    
    // Test data
    req := &CreateEntityRequest{
        Name:  "Test Entity",
        Email: "test@example.com",
    }
    
    // Mock repository call
    mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*models.YourEntity")).Return(nil)
    
    // Execute test
    result, err := service.Create(context.Background(), req)
    
    // Assertions
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, req.Name, result.Name)
    assert.Equal(t, req.Email, result.Email)
    
    // Verify mock calls
    mockRepo.AssertExpectations(t)
}
```

### Integration Testing
- ทดสอบการทำงานร่วมกันระหว่าง components
- ใช้ test databases สำหรับการทดสอบ
- ทดสอบ API endpoints ด้วย httptest

```go
// Integration test example
func TestController_Create(t *testing.T) {
    // Setup test database
    db := setupTestDB(t)
    defer cleanupTestDB(t, db)
    
    // Setup dependencies
    repo := repositories.NewRepository(db)
    service := services.NewService(repo)
    controller := controllers.NewController(service)
    
    // Setup Gin router
    router := gin.New()
    router.POST("/entities", controller.Create)
    
    // Test data
    reqBody := `{"name": "Test Entity", "email": "test@example.com"}`
    req, _ := http.NewRequest("POST", "/entities", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    
    // Execute request
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assertions
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response responses.SuccessResponse
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.True(t, response.Success)
}
```

### Database Testing
- ใช้ Docker containers สำหรับ test databases
- ใช้ database migrations สำหรับการจัดการ schema
- ทำ cleanup หลังการทดสอบเสมอ

```go
// Test database setup
func setupTestDB(t *testing.T) *sqlx.DB {
    // Use test database
    dsn := "postgres://test:test@localhost:5432/test_db?sslmode=disable"
    db, err := sqlx.Connect("postgres", dsn)
    require.NoError(t, err)
    
    // Run migrations
    err = runMigrations(db)
    require.NoError(t, err)
    
    return db
}

func cleanupTestDB(t *testing.T, db *sqlx.DB) {
    // Clean up test data
    db.Exec("TRUNCATE TABLE entities CASCADE")
    db.Close()
}
```

### Performance Testing
- ใช้ Go's benchmark testing
- ทดสอบ API performance ด้วย tools เช่น wrk, ab
- ตรวจสอบ memory usage และ goroutine leaks

```go
// Benchmark test example
func BenchmarkService_Create(b *testing.B) {
    service := setupService()
    req := &CreateEntityRequest{
        Name:  "Benchmark Entity",
        Email: "benchmark@example.com",
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := service.Create(context.Background(), req)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

## 🔄 CI/CD

โปรเจคนี้ใช้ GitHub Actions สำหรับการทำ CI/CD:

### Continuous Integration
- **Code Quality**: ใช้ golangci-lint สำหรับ static analysis
- **Testing**: รัน unit tests และ integration tests
- **Security**: ใช้ gosec สำหรับ security scanning
- **Build**: สร้าง binary สำหรับหลาย platforms

### Continuous Deployment
- **Development**: Deploy อัตโนมัติเมื่อ merge เข้า `develop` branch
- **Staging**: Deploy อัตโนมัติเมื่อ merge เข้า `main` branch
- **Production**: Deploy ด้วย manual approval หรือ semantic versioning

### Docker Integration
- **Multi-stage builds**: ลดขนาด image ใน production
- **Security scanning**: ใช้ Trivy สำหรับ vulnerability scanning
- **Registry**: Push ไปยัง container registry

```yaml
# Example GitHub Actions workflow
name: CI/CD Pipeline

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:13
        env:
          POSTGRES_PASSWORD: test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.23.3
    
    - name: Cache Go modules
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
    
    - name: Install dependencies
      run: go mod download
    
    - name: Run tests
      run: go test -v -race -coverprofile=coverage.out ./...
    
    - name: Upload coverage
      uses: codecov/codecov-action@v3
      with:
        file: ./coverage.out
    
    - name: Run linter
      uses: golangci/golangci-lint-action@v3
      with:
        version: latest
  
  build:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.23.3
    
    - name: Build binary
      run: CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o microservice main.go
    
    - name: Build Docker image
      run: docker build -t microservice:${{ github.sha }} .
    
    - name: Push to registry
      run: |
        echo ${{ secrets.DOCKER_PASSWORD }} | docker login -u ${{ secrets.DOCKER_USERNAME }} --password-stdin
        docker push microservice:${{ github.sha }}
```

## 📘 เอกสารอ้างอิงเพิ่มเติม

- [Go Documentation](https://golang.org/doc/)
- [Gin Web Framework](https://gin-gonic.com/docs/)
- [Go Database Drivers](https://github.com/golang/go/wiki/SQLDrivers)
- [JWT in Go](https://github.com/golang-jwt/jwt)
- [Zap Logging](https://github.com/uber-go/zap)
- [Go Testing](https://golang.org/pkg/testing/)
- [Docker for Go](https://docs.docker.com/language/go/)
- [Go Performance Best Practices](https://go.dev/doc/effective_go)
