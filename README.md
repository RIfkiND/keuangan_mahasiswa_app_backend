# GoRestAPI

A RESTful API built with Go, GORM, and PostgreSQL, ready for Docker deployment.

## Features

- Modular project structure (handlers, services, repositories, models)
- GORM ORM for PostgreSQL
- Authentication endpoints
- CRUD for Posts, Categories, Post Images
- Environment variable support via `.env`
- API versioning (`/v1`)
- Dockerized for easy deployment

## Getting Started

### Prerequisites

- [Go 1.23+](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- PostgreSQL database (e.g., [Supabase](https://supabase.com/))

### Environment Variables

Create a `.env` file in the project root:

```env
SUPABASE_URL=your-supabase-url
SUPABASE_ANON_KEY=your-anon-key
SUPABASE_DB_URL=postgres://user:password@host:port/dbname
JWT_SECRET=your-jwt-secret
```

### Running Locally

```bash
go run cmd/main.go
```

### Database Migration

To auto-migrate your database tables:

```bash
go run cmd/migration/migration.go
```

### Docker Usage

**Build the Docker image:**
```bash
docker build -t gorestapi .
```

**Run the container:**
```bash
docker run --env-file .env -p 8080:8080 gorestapi
```

## API Endpoints

### Authentication

| Method | Endpoint             | Description         |
|--------|----------------------|---------------------|
| POST   | `/v1/auth/register`  | Register user       |
| POST   | `/v1/auth/login`     | Login user          |

### Posts

| Method | Endpoint                 | Description         |
|--------|--------------------------|---------------------|
| POST   | `/v1/posts/create`       | Create a post       |
| GET    | `/v1/posts/get?id=1`     | Get post by ID      |
| GET    | `/v1/posts/all`          | List all posts      |
| PUT    | `/v1/posts/update?id=1`  | Update a post       |
| DELETE | `/v1/posts/delete?id=1`  | Delete a post       |

### Categories

| Method | Endpoint                      | Description         |
|--------|-------------------------------|---------------------|
| POST   | `/v1/categories/create`       | Create a category   |
| GET    | `/v1/categories/all`          | List all categories |
| PUT    | `/v1/categories/update?id=1`  | Update a category   |
| DELETE | `/v1/categories/delete?id=1`  | Delete a category   |

### Post Images

| Method | Endpoint                        | Description         |
|--------|---------------------------------|---------------------|
| POST   | `/v1/postimages/create`         | Add image to post   |
| GET    | `/v1/postimages/all`            | List all images     |
| PUT    | `/v1/postimages/update?id=1`    | Update image        |
| DELETE | `/v1/postimages/delete?id=1`    | Delete image        |

## Project Structure

```
.
├── cmd/                # Entry points (main.go, migration.go)
├── config/             # Database and config setup
├── internals/
│   ├── handlers/       # HTTP handlers (v1/)
│   ├── models/         # GORM models (User, Post, Category, PostImage)
│   ├── repositories/   # Data access layer
│   └── services/       # Business logic
├── go.mod
├── dockerfile
└── .env
```

## License

**Happy coding! 🚀**