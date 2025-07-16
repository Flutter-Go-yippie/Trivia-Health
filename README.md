# TriviaHealth - AI-Powered Fitness Assistant

![TriviaHealth Banner](apps/android-app/assets/robot.png)

TriviaHealth is a comprehensive AI-powered fitness application that provides personalized workout plans, real-time coaching, and progress tracking. The platform combines cutting-edge AI technology with user-friendly mobile and web interfaces to deliver a complete fitness solution.

## 🏗️ Architecture Overview

TriviaHealth follows a modern microservices architecture with the following components:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Flutter App    │    │   REST API      │    │   AI Services   │
│  (Mobile/Web)   │◄──►│   (Go/Gin)      │◄──►│  (OpenRouter)   │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │   Databases     │
                    │ PostgreSQL +    │
                    │   MongoDB       │
                    └─────────────────┘
```

## 🚀 Features

### 🎯 Core Features
- **AI-Powered Workout Generation**: Personalized workout plans based on user profile, goals, and preferences
- **Smart Chat Assistant**: Interactive AI coach for fitness guidance and motivation
- **Progressive Training**: Adaptive workout difficulty based on user progress
- **Real-Time Exercise Tracking**: Live exercise guidance with rest timers and form tips
- **Progress Analytics**: Comprehensive progress tracking with streaks and achievements
- **Multi-Platform Support**: Cross-platform mobile app (iOS, Android, Web, Desktop)

### 🤖 AI Integration
- **Multiple AI Models**: Integration with 8+ AI models through OpenRouter API
- **Automatic Model Switching**: Smart fallback system for reliability
- **Contextual Understanding**: Chat history preservation for better assistance
- **Structured Response Processing**: JSON-based workout plan generation

### 💪 Workout Management
- **Personalized Plans**: Custom workout plans based on:
  - Age, height, weight
  - Fitness goals (weight loss, muscle gain, endurance, etc.)
  - Fitness level (beginner, intermediate, advanced)
  - Available time and health considerations
- **Exercise Library**: Comprehensive exercise database with:
  - Detailed instructions and form tips
  - Target muscle groups
  - Set/rep recommendations
  - Rest period optimization
- **Plan Regeneration**: AI-powered plan updates based on user feedback

### 📊 Progress Tracking
- **Workout Completion**: Track completed exercises and workouts
- **Streak Monitoring**: Consecutive day tracking for motivation
- **Performance Analytics**: Progress visualization and insights
- **Achievement System**: Level progression and milestone rewards

## 🛠️ Technology Stack

### Frontend (Mobile App)
- **Framework**: Flutter 3.8.1+
- **State Management**: BLoC (Business Logic Component)
- **Navigation**: Auto Route 7.8.4
- **HTTP Client**: Dio with smart retry mechanisms
- **Local Storage**: SharedPreferences
- **UI Components**: Custom design system with Material Design

### Backend (REST API)
- **Language**: Go 1.24.4
- **Web Framework**: Gorilla Mux
- **Authentication**: JWT (JSON Web Tokens)
- **Documentation**: Swagger/OpenAPI 3.0
- **Validation**: Go Playground Validator
- **Migration**: golang-migrate

### Databases
- **PostgreSQL**: Primary database for user data, profiles, and authentication
- **MongoDB**: Workout plans, AI responses, and chat history storage

### AI & External Services
- **OpenRouter API**: Multi-model AI service integration
- **Supported Models**:
  - DeepSeek Chat v3
  - Google Gemma 7B
  - Claude 3 Haiku
  - LLaMA 3 8B
  - Microsoft Phi-3 Mini
  - And more...

### DevOps & Deployment
- **Containerization**: Docker & Docker Compose
- **Environment**: Production-ready configuration
- **Networking**: Internal Docker networking for security

## 📱 Mobile App Structure

```
apps/android-app/
├── lib/
│   ├── app/                    # App configuration & routing
│   │   ├── app_router.dart     # Navigation setup
│   │   ├── app_dependencies.dart # Dependency injection
│   │   └── data/               # Network services
│   ├── constants/              # Design system
│   │   ├── app_colors.dart     # Color palette
│   │   └── app_text_styles.dart # Typography
│   ├── features/               # Feature modules
│   │   ├── auth/               # Authentication flow
│   │   ├── chat/               # AI chat assistant
│   │   ├── home/               # Main navigation
│   │   ├── path/               # Workout tracking
│   │   ├── settings/           # User preferences
│   │   └── init/               # App initialization
│   └── uikit/                  # Reusable UI components
├── test/                       # Unit & widget tests
└── releases/                   # APK builds
```

### Key App Features
- **Multi-step Registration**: Comprehensive user onboarding with profile setup
- **Biometric Authentication**: JWT token management with remember me functionality
- **Offline Support**: Local data caching and sync capabilities
- **Real-time Chat**: Interactive AI assistant with chat history
- **Exercise Execution**: Live workout tracking with timers and guidance
- **Progress Visualization**: Charts and statistics for user progress

## 🌐 REST API Structure

```
apps/rest-api/
├── cmd/server/                 # Application entry point
├── internal/
│   ├── config/                 # Configuration management
│   ├── handlers/               # HTTP request handlers
│   │   ├── auth.go            # Authentication endpoints
│   │   ├── chat.go            # Chat functionality
│   │   ├── profile.go         # User profile management
│   │   └── workout.go         # Workout plan endpoints
│   ├── middleware/            # HTTP middleware
│   │   ├── auth.go            # JWT authentication
│   │   ├── logging.go         # Request logging
│   │   └── validation.go      # Input validation
│   ├── models/                # Data models
│   │   ├── auth.go            # Authentication models
│   │   ├── fitness.go         # Workout & exercise models
│   │   └── progress.go        # Progress tracking models
│   ├── repository/            # Data access layer
│   │   ├── postgres.go        # PostgreSQL operations
│   │   └── mongodb.go         # MongoDB operations
│   └── services/              # Business logic
│       ├── auth.go            # Authentication service
│       ├── ai.go              # AI integration service
│       ├── profile.go         # Profile management
│       └── openrouter.go      # AI model client
├── migrations/                # Database migrations
├── docs/                      # API documentation
└── pkg/utils/                 # Shared utilities
```

## 🔐 API Endpoints

### Authentication
```http
POST /register                 # User registration
POST /login                   # User authentication
```

### User Profile
```http
POST /api/profile             # Save/update fitness profile
GET  /api/profile             # Get user profile
```

### Workout Management
```http
POST /api/generate-plan       # Generate new workout plan
GET  /api/workout-plan        # Get current workout plan
POST /api/regenerate-plan     # Update plan with feedback
POST /api/complete-workout    # Mark workout as completed
```

### AI Chat & Analytics
```http
POST /api/chat                # Chat with AI assistant
GET  /api/chat/history        # Get chat history
GET  /api/progress            # Get user progress
GET  /api/rating              # Get user rating/leaderboard
GET  /api/motivation          # Get motivational messages
```

### System
```http
GET  /health                  # Health check endpoint
GET  /swagger/                # API documentation
```

## 💾 Database Schema

### PostgreSQL Schema
```sql
-- Users and authentication
users (id, email, password_hash, created_at)

-- Fitness profiles
fitness_profiles (id, user_id, height_cm, weight_kg, age, 
                 fitness_goal, timeframe, fitness_level, 
                 weekly_time_minutes, updated_at)

-- Health considerations
health_issues (id, name)
user_health_issues (user_id, issue_id)

-- Chat history
chat_messages (id, user_id, message, response, is_user, created_at)
```

### MongoDB Collections
```javascript
// Workout plans with full scheduling
workout_plans {
  user_id: int,
  title: string,
  workouts: [
    {
      workout_id: ObjectId,
      name: string,
      description: string,
      status: "planned|current|completed",
      scheduled_date: Date,
      exercises: [
        {
          exercise_id: ObjectId,
          name: string,
          muscle_group: string,
          sets: int,
          reps: int,
          rest_sec: int,
          notes: string,
          technique: string
        }
      ]
    }
  ],
  status: boolean,
  created_at: Date,
  updated_at: Date
}

// User progress tracking
user_progress {
  user_id: int,
  total_workouts: int,
  consecutive_days: int,
  level: string,
  completed_workouts: [string],
  last_workout_date: Date
}

// Workout completion records
workout_completions {
  user_id: int,
  workout_id: ObjectId,
  completed_at: Date
}
```

## 🧪 Testing

The project includes comprehensive testing coverage:

### Backend Testing (22 test files)
- **Handlers**: HTTP endpoint testing (6 files)
- **Middleware**: Authentication and validation testing (3 files)  
- **Services**: Business logic testing (6 files)
- **Repository**: Data layer testing (2 files)
- **Utils**: JWT and crypto utilities testing (3 files)
- **Models**: Data validation testing (2 files)

### Frontend Testing
- **Unit Tests**: BLoC state management testing
- **Widget Tests**: UI component testing
- **Integration Tests**: End-to-end user flow testing

### Test Coverage Areas
- ✅ Authentication & JWT validation
- ✅ Workout plan generation & AI integration
- ✅ User profile management
- ✅ Progress tracking algorithms
- ✅ API endpoint validation
- ✅ Database operations
- ✅ Error handling & edge cases

## 🐳 Docker Deployment

The application includes a complete Docker Compose setup:

```yaml
# docker-compose.yml
services:
  rest_api:          # Go REST API server
  postgres:          # PostgreSQL database
  mongo:             # MongoDB database
```

### Environment Configuration
```bash
# Required environment variables
OPENROUTER_KEY=sk-or-v1-your-key-here
DATABASE_URL=postgres://user:pass@host:port/db
MONGOURI=mongodb://host:port/database
PORT=80
```

## 🚀 Getting Started

### Prerequisites
- **Go 1.24.4+** for backend development
- **Flutter 3.8.1+** for mobile app development
- **Docker & Docker Compose** for deployment
- **PostgreSQL 15+** for primary database
- **MongoDB 4.4+** for document storage
- **OpenRouter API Key** for AI integration

### Local Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/Flutter-Go-yippie/Trivia-Health.git
   cd Trivia-Health
   ```

2. **Setup environment variables**
   ```bash
   # Create .env file in apps/rest-api/
   OPENROUTER_KEY=your-openrouter-api-key
   DATABASE_URL=postgres://postgres:postgres@localhost:5432/fitness_ai?sslmode=disable
   MONGOURI=mongodb://localhost:27017/fitness_ai
   JWT_SECRET=your-jwt-secret
   PORT=8080
   ```

3. **Start with Docker Compose**
   ```bash
   docker-compose up -d
   ```

4. **Run database migrations**
   ```bash
   cd apps/rest-api
   make migrate-up
   ```

5. **Start the backend server**
   ```bash
   cd apps/rest-api
   go mod download
   go run cmd/server/main.go
   ```

6. **Setup Flutter app**
   ```bash
   cd apps/android-app
   flutter pub get
   flutter run
   ```

### Building for Production

#### Backend
```bash
cd apps/rest-api
go build -o server cmd/server/main.go
```

#### Mobile App
```bash
cd apps/android-app
flutter build apk --release
# or
flutter build web --release
```

## 📊 AI Model Integration

TriviaHealth integrates with multiple AI models through OpenRouter API for enhanced reliability:

### Available Models
1. **DeepSeek Chat v3** - Primary model for best performance
2. **Google Gemma 7B** - Backup for general queries  
3. **Claude 3 Haiku** - Conversational AI excellence
4. **LLaMA 3 8B** - Meta's instruction-tuned model
5. **Microsoft Phi-3 Mini** - Lightweight but powerful
6. **OpenChat 7B** - Optimized for conversations
7. **Nous Hermes Mixtral** - Advanced reasoning
8. **Cypher Alpha** - Alternative reasoning model

### AI Features
- **Automatic Model Switching**: Failover to alternative models
- **Smart Retry System**: Multiple attempts with different models
- **Context Preservation**: Chat history for better understanding
- **Structured Responses**: JSON workout plan generation
- **Cost Optimization**: Free tier model usage prioritization

## 🔒 Security Features

### Authentication & Authorization
- **JWT Token Authentication**: Secure stateless authentication
- **Password Hashing**: bcrypt for secure password storage
- **Token Expiration**: Configurable token lifetime
- **Request Validation**: Input sanitization and validation
- **CORS Protection**: Cross-origin request security

### Data Protection
- **Environment Variables**: Sensitive data protection
- **Database Encryption**: Secure data storage
- **API Rate Limiting**: Protection against abuse
- **Input Validation**: SQL injection prevention
- **Error Handling**: Secure error responses

## 🚀 Performance Features

### Backend Optimization
- **Connection Pooling**: PostgreSQL and MongoDB connection management
- **Caching Strategy**: Efficient data caching
- **Graceful Shutdown**: Clean server termination
- **Timeout Management**: Request timeout handling
- **Concurrent Processing**: Go's goroutine utilization

### Mobile App Optimization
- **State Management**: Efficient BLoC pattern implementation
- **Image Caching**: Optimized image loading
- **Offline Support**: Local data persistence
- **Lazy Loading**: On-demand resource loading
- **Memory Management**: Efficient resource cleanup

## 📈 Monitoring & Analytics

### Health Monitoring
- **Health Check Endpoints**: System status monitoring
- **Database Connection Testing**: Real-time connectivity checks
- **API Response Time Tracking**: Performance monitoring
- **Error Rate Monitoring**: System reliability tracking

### User Analytics
- **Workout Completion Rates**: User engagement metrics
- **Progress Tracking**: Fitness goal achievement
- **Feature Usage**: App functionality analytics
- **User Retention**: Long-term engagement analysis

## 🤝 Contributing

We welcome contributions to TriviaHealth! Please follow these guidelines:

### Development Workflow
1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** with proper testing
4. **Run tests**: `go test ./...` and `flutter test`
5. **Commit changes**: `git commit -m 'Add amazing feature'`
6. **Push to branch**: `git push origin feature/amazing-feature`
7. **Open a Pull Request**

### Code Standards
- **Go**: Follow effective Go practices and gofmt formatting
- **Flutter**: Adhere to Dart style guide and widget conventions
- **Testing**: Maintain or improve test coverage
- **Documentation**: Update relevant documentation
- **Commit Messages**: Use conventional commit format

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## 👥 Team

TriviaHealth is developed by the Flutter-Go-yippie team, combining expertise in:
- **Flutter/Dart Mobile Development**
- **Go Backend Development**  
- **AI/ML Integration**
- **Database Design**
- **DevOps & Cloud Infrastructure**

## 🆘 Support

For support, bug reports, or feature requests:

- **Issues**: [GitHub Issues](https://github.com/Flutter-Go-yippie/Trivia-Health/issues)
- **Discussions**: [GitHub Discussions](https://github.com/Flutter-Go-yippie/Trivia-Health/discussions)
- **Documentation**: [API Docs](http://localhost:8080/swagger/)

## 🚧 Roadmap

### Upcoming Features
- [ ] **Social Features**: Friend challenges and leaderboards
- [ ] **Nutrition Tracking**: Meal planning and calorie tracking
- [ ] **Wearable Integration**: Smartwatch and fitness tracker sync
- [ ] **Video Workouts**: AI-powered form analysis
- [ ] **Community Features**: User-generated content and tips
- [ ] **Advanced Analytics**: ML-powered insights and predictions

### Technical Improvements
- [ ] **GraphQL API**: Enhanced query capabilities
- [ ] **Real-time Updates**: WebSocket integration
- [ ] **Offline Mode**: Complete offline workout capability
- [ ] **Performance Optimization**: Further speed improvements
- [ ] **Internationalization**: Multi-language support
- [ ] **Advanced Testing**: E2E automation and load testing

---

**TriviaHealth** - Your AI-Powered Personal Fitness Assistant 🤖💪

*Built with ❤️ using Flutter, Go, and cutting-edge AI technology*
