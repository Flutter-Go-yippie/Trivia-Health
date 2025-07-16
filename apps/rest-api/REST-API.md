# TriviaHealth REST API Documentation

## Overview
REST API for the TriviaHealth fitness application providing AI-powered workout planning, chat functionality, and user progress tracking.

## Base URL
```
http://localhost:8080
```

## Authentication
All protected endpoints require Bearer token authentication:
```
Authorization: Bearer <access_token>
```

## Quick Start

### 1. Start Services
```bash
make db-up mongo-up    # Start databases
make migrate-up        # Run migrations
make run-api          # Start API server
```

### 2. Register User
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

### 3. Access Swagger UI
```
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Authentication

#### Register User
```http
POST /register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Login
```http
POST /login
Content-Type: application/json

{
  "email": "user@example.com", 
  "password": "password123"
}
```

#### Refresh Token
```http
POST /refresh
Content-Type: application/json

{
  "refresh_token": "your_refresh_token"
}
```

### Profile Management

#### Save Fitness Profile
```http
POST /api/profile
Authorization: Bearer <token>
Content-Type: application/json

{
  "height": 175.0,
  "weight": 70.0,
  "age": 25,
  "goal": "weight_loss",
  "fitness_level": "intermediate",
  "timeframe": "3months",
  "available_minutes": 180,
  "health_issues": ["knee_pain"]
}
```

#### Get Profile
```http
GET /api/profile
Authorization: Bearer <token>
```

### Workout Planning

#### Generate Workout Plan
```http
POST /api/generate-plan
Authorization: Bearer <token>
```

#### Get Current Plan
```http
GET /api/workout-plan
Authorization: Bearer <token>
```

#### Regenerate Plan
```http
POST /api/regenerate-plan
Authorization: Bearer <token>
Content-Type: application/json

{
  "comments": "Make it more challenging with more cardio"
}
```

#### Complete Workout
```http
POST /api/complete-workout
Authorization: Bearer <token>
Content-Type: application/json

{
  "workout_id": "workout_object_id"
}
```

### Progress Tracking

#### Get User Progress
```http
GET /api/progress
Authorization: Bearer <token>
```

#### Get User Rating/Leaderboard
```http
GET /api/rating
Authorization: Bearer <token>
```

### AI Chat

#### Send Message
```http
POST /api/chat
Authorization: Bearer <token>
Content-Type: application/json

{
  "message": "How do I improve my squat form?"
}
```

#### Get Chat History
```http
GET /api/chat/history
Authorization: Bearer <token>
```

#### Get Motivational Message
```http
GET /api/motivation
Authorization: Bearer <token>
```

### Exercise Media

#### Save Exercise Media
```http
POST /api/exercise/media
Authorization: Bearer <token>
Content-Type: application/json

{
  "exercise_id": "exercise_object_id",
  "image_url": "https://example.com/image.jpg",
  "description": "Proper squat form",
  "order": 1
}
```

#### Get Exercise Media
```http
GET /api/exercise/{exercise_id}/media
Authorization: Bearer <token>
```

#### Delete Exercise Media
```http
DELETE /api/exercise/media/{media_id}
Authorization: Bearer <token>
```

### Health Check
```http
GET /health
```

## Data Models

### User Registration/Login
```json
{
  "email": "string",
  "password": "string (min 8 chars)"
}
```

### Auth Response
```json
{
  "access_token": "string",
  "refresh_token": "string", 
  "token_type": "Bearer",
  "expires_in": 900,
  "email": "string"
}
```

### Fitness Profile
```json
{
  "height": 175.0,
  "weight": 70.0,
  "age": 25,
  "goal": "weight_loss|muscle_gain|endurance|flexibility|general_fitness",
  "fitness_level": "beginner|intermediate|advanced",
  "timeframe": "1month|3months|6months|1year",
  "available_minutes": 180,
  "health_issues": ["string"]
}
```

### Workout Plan
```json
{
  "id": "object_id",
  "user_id": 1,
  "title": "string",
  "status": true,
  "created_at": "2024-01-01T00:00:00Z",
  "workouts": [
    {
      "workout_id": "object_id",
      "name": "Upper Body Strength",
      "description": "Focus on chest, back, shoulders",
      "status": "planned|done|expired",
      "scheduled_date": "2024-01-01T00:00:00Z",
      "exercises": [
        {
          "exercise_id": "object_id",
          "name": "Push-ups",
          "muscle_group": "Chest",
          "sets": 3,
          "reps": 12,
          "rest_sec": 60,
          "notes": "Keep core tight",
          "technique": "Slow and controlled"
        }
      ]
    }
  ]
}
```

### User Progress
```json
{
  "user_id": 1,
  "total_workouts": 15,
  "consecutive_days": 5,
  "level": "Intermediate",
  "completed_workouts": ["Workout 1", "Workout 2"],
  "last_workout_date": "2024-01-01T00:00:00Z"
}
```

## Error Responses
```json
{
  "error": "Bad Request",
  "message": "Detailed error description"
}
```

## Status Codes
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `409` - Conflict
- `500` - Internal Server Error
- `503` - Service Unavailable

## Environment Variables
```bash
# Database
DATABASE_URL=postgres://postgres:postgres@localhost:5432/fitness_ai?sslmode=disable
MONGOURI=mongodb://localhost:27017/fitness_ai
MONGODBNAME=fitness_ai

# JWT
JWT_SECRET=your-secret-key
JWT_EXPIRATION=15m
REFRESH_EXPIRATION=7d

# AI Service
OPENROUTER_KEY=sk-or-v1-your-key-here

# Server
PORT=8080
ENVIRONMENT=development
```

## Development Commands
```bash
# Database
make db-up          # Start PostgreSQL
make mongo-up       # Start MongoDB
make migrate-up     # Run migrations

# Development
make run-api        # Start API server
make test          # Run all tests
make test-coverage # Run tests with coverage

# Build
make build         # Build binary
make swagger       # Generate docs
```

## Testing
- Unit tests: 22 test files covering handlers, services, middleware
- Integration tests: API endpoint testing
- Coverage: High coverage of critical business logic
- Benchmarks: Performance testing for algorithms

## AI Features
- **Models**: 8 free AI models with automatic switching
- **Chat**: Contextual fitness conversations
- **Plans**: Personalized workout generation
- **Beginner Mode**: Simplified explanations for beginners
- **Motivation**: AI-generated motivational messages

## Database Schema
- **PostgreSQL**: Users, profiles, health issues
- **MongoDB**: Workouts, chat history, progress, media

## Architecture
- Clean architecture with separated layers
- Repository pattern for data access
- Service layer for business logic
- Middleware for cross-cutting concerns
- Swagger documentation generation