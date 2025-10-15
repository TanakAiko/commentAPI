# Comment API

A RESTful microservice API for managing comments in a real-time forum application. Built with Go and SQLite, this service handles comment creation, retrieval, updates, and deletion with support for likes and dislikes.

<div align="center">

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![REST API](https://img.shields.io/badge/REST-API-green?style=for-the-badge)

</div>

## ✨ Features

- **Create Comments**: Add new comments to posts
- **Retrieve Comments**: Get all comments for a specific post or fetch the latest comment
- **Update Likes/Dislikes**: Manage user reactions to comments
- **Delete Comments**: Remove comments from the database
- **SQLite Database**: Lightweight and persistent storage
- **Docker Support**: Easy deployment with containerization

## 📋 Prerequisites

- Go 1.20 or higher
- SQLite3
- Docker (optional, for containerized deployment)

## 🛠️ Installation

### Local Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd commentAPI
```

2. Install dependencies:
```bash
go mod download
```

3. Build and run the application:
```bash
go build -o commentapi-server
./commentapi-server
```

Or simply run directly:
```bash
go run main.go
```

The server will start on `http://localhost:8084`

### Docker Setup

1. Build the Docker image:
```bash
docker build -t comment-api .
```

2. Run the container:
```bash
docker run -p 8084:8084 -v $(pwd)/databases:/app/databases comment-api
```

## 📡 API Endpoints

The API uses a single endpoint with different actions specified in the request body.

### Endpoint
```
POST /
```

### Request Format
All requests should be sent as JSON with the following structure:
```json
{
  "action": "<action_name>",
  "body": {
    // Comment object fields
  }
}
```

### Available Actions

#### 1. Create Comment
```json
{
  "action": "createComment",
  "body": {
    "postID": 1,
    "userID": 123,
    "nickname": "john_doe",
    "content": "This is a great post!",
    "nbrLike": 0,
    "nbrDislike": 0
  }
}
```

#### 2. Get All Comments for a Post
```json
{
  "action": "getAllPostComment",
  "body": {
    "postID": 1
  }
}
```

#### 3. Get Last Comment
```json
{
  "action": "getLastComment",
  "body": {}
}
```

#### 4. Update Like/Dislike
```json
{
  "action": "updateLike",
  "body": {
    "commentID": 1,
    "userID": 123,
    "nickname": "john_doe",
    "likedBy": ["user1", "user2"],
    "dislikedBy": [],
    "nbrLike": 2,
    "nbrDislike": 0
  }
}
```

#### 5. Delete Comment
```json
{
  "action": "delete",
  "body": {
    "commentID": 1
  }
}
```

## 📁 Project Structure

```
commentAPI/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── Dockerfile              # Docker configuration
├── README.md              # This file
├── config/
│   └── constants.go       # Configuration constants (port, etc.)
├── databases/
│   └── sqlRequests/
│       ├── createTable.sql         # Table schema
│       └── insertNewComment.sql    # Insert query template
├── internals/
│   ├── dbManager/
│   │   └── initDB.go      # Database initialization
│   ├── handlers/
│   │   ├── mainHandler.go  # Main request router
│   │   ├── createHandler.go # Create comment handler
│   │   ├── getHandler.go    # Get comments handler
│   │   ├── deleteHandler.go # Delete comment handler
│   │   └── update.go        # Update like/dislike handler
│   └── tools/
│       └── utils.go        # Utility functions
├── models/
│   ├── comment.go         # Comment model and methods
│   └── request.go         # Request structure
└── scripts/
    ├── init.sh            # Repository clone helper script
    └── push.sh            # Deployment script
```

## 💾 Database Schema

The application uses SQLite with the following schema:

```sql
CREATE TABLE IF NOT EXISTS comments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    postId INTEGER NOT NULL,
    userId INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    likedBy TEXT NOT NULL,
    dislikedBy TEXT NOT NULL,
    content TEXT NOT NULL,
    nbrLike INTEGER,
    nbrDislike INTEGER,
    createdAt DATETIME NOT NULL
);
```

## 🔧 Configuration

The API configuration can be modified in `config/constants.go`:

- **Port**: Default is `8084`

## 🧪 Testing

You can test the API using curl:

```bash
# Create a comment
curl -X POST http://localhost:8084/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "createComment",
    "body": {
      "postID": 1,
      "userID": 123,
      "nickname": "test_user",
      "content": "Test comment",
      "nbrLike": 0,
      "nbrDislike": 0
    }
  }'

# Get all comments for a post
curl -X POST http://localhost:8084/ \
  -H "Content-Type: application/json" \
  -d '{
    "action": "getAllPostComment",
    "body": {
      "postID": 1
    }
  }'
```

## 📦 Dependencies

- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for Go

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🔗 Related Projects

This microservice is designed to work as part of a larger real-time forum ecosystem. Make sure to set up the related services for full functionality.

---

<div align="center">

**⭐ Star this repository if you found it helpful! ⭐**

Made with ❤️ from 🇸🇳

</div>