# TriviaHealth - AI-Powered Fitness Assistant

![TriviaHealth Banner](apps/android-app/assets/robot.png)

TriviaHealth is an AI-powered fitness application that provides personalized workout plans, coaching, and progress tracking using Flutter and Go.

## 🏗️ Architecture

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│  Flutter App    │    │   REST API      │    │   AI Services   │
│  (Mobile/Web)   │◄──►│   (Go)          │◄──►│  (OpenRouter)   │
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

- **AI Workout Generation**: Personalized plans based on user goals
- **Chat Assistant**: Interactive AI coach for guidance
- **Progress Tracking**: Workout completion and streak monitoring  
- **Cross-Platform**: Flutter app for mobile and web
- **Real-Time Guidance**: Exercise tracking with timers

## 🛠️ Tech Stack

- **Frontend**: Flutter, BLoC state management
- **Backend**: Go, REST API, JWT authentication
- **Databases**: PostgreSQL + MongoDB
- **AI**: OpenRouter API with multiple models
- **DevOps**: Docker, CI/CD pipelines

## � Quick Start

1. **Clone and setup**
   ```bash
   git clone https://github.com/Flutter-Go-yippie/Trivia-Health.git
   cd Trivia-Health
   ```

2. **Start with Docker**
   ```bash
   docker-compose up -d
   ```

3. **Run Flutter app**
   ```bash
   cd apps/android-app
   flutter pub get && flutter run
   ```

## � API Endpoints

- `POST /register` - User registration
- `POST /login` - Authentication  
- `POST /api/generate-plan` - Generate workout plan
- `POST /api/chat` - AI chat assistant
- `GET /api/progress` - User progress
- `GET /swagger/` - API documentation

## 📚 Documentation

### Project Documentation
- **[REST API Documentation](apps/rest-api/REST-API.md)** - Complete API endpoints and usage guide
- **[AI Models Guide](apps/rest-api/AI_MODELS.md)** - AI integration and model configuration
- **[Testing Guide](apps/rest-api/README_TESTING.md)** - Backend testing setup and guidelines
- **[Flutter App Guide](apps/android-app/README.md)** - Mobile app setup and development

### Additional Resources
- **[Swagger API Docs](http://localhost:8080/swagger/index.html)** - Interactive API documentation
- **[Project Architecture](apps/rest-api/REST-API.md#architecture)** - Detailed system design
- **[Database Schema](apps/rest-api/REST-API.md#database)** - Database structure and relationships

## Implementation checklist

### Technical requirements (20 points)
#### Backend development (8 points)
- [X] Go-based backend (3 points)
- [X] RESTful API with Swagger documentation (2 point)
- [X] PostgreSQL database with proper schema design (1 point)
- [X] JWT-based authentication and authorization (1 point)
- [X] Comprehensive unit and integration tests (1 point)

#### Frontend development (8 points)
- [x] Flutter-based cross-platform application (mobile + web) (3 points)
- [x] Responsive UI design with custom widgets (1 point)
- [x] State management implementation (1 point)
- [ ] Offline data persistence (1 point)
- [x] Unit and widget tests (1 point)
- [ ] Support light and dark mode (1 point)

#### DevOps & deployment (4 points)
- [x] Docker compose for all services (1 point)
- [x] CI/CD pipeline implementation (1 point)
- [x] Environment configuration management using config files (1 point)
- [x] GitHub pages for the project (1 point)

### Non-Technical Requirements (10 points)
#### Project management (4 points)
- [x] GitHub organization with well-maintained repository (1 point)
- [x] Regular commits and meaningful pull requests from all team members (1 point)
- [x] Project board (GitHub Projects) with task tracking (1 point)
- [x] Team member roles and responsibilities documentation (1 point)

#### Documentation (4 points)
- [x] Project overview and setup instructions (1 point)
- [x] Screenshots and GIFs of key features (1 point)
- [x] API documentation (1 point)
- [x] Architecture diagrams and explanations (1 point)

#### Code quality (2 points)
- [x] Consistent code style and formatting during CI/CD pipeline (1 point)
- [x] Code review participation and resolution (1 point)

### Bonus Features (up to 10 points)
- [ ] Localization for Russian (RU) and English (ENG) languages (2 points)
- [x] Good UI/UX design (up to 3 points)
- [x] Integration with external APIs (fitness trackers, health devices) (up to 5 points)
- [x] Comprehensive error handling and user feedback (up to 2 points)
- [ ] Advanced animations and transitions (up to 3 points)
- [x] Widget implementation for native mobile elements (up to 2 points)
- [x] MongoDB non-relational database usage for workout data (2 points)

Total points implemented: 28/30 (excluding bonus points)

---

**TriviaHealth** - Your AI-Powered Personal Fitness Assistant 🤖💪
