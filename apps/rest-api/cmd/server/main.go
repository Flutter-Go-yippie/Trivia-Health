// @title TriviaHealth API
// @version 1.0
// @description REST API for the TriviaHealth fitness application
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token authentication. Format: Bearer {token}
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "rest-api/docs"
	"rest-api/internal/config"
	"rest-api/internal/handlers"
	"rest-api/internal/middleware"
	"rest-api/internal/repository"
	"rest-api/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	fmt.Printf("Loaded config: %+v\n", cfg)

	// Run database migrations
	if err := runMigrations(cfg.DatabaseURL); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	// Initialize repository
	postgresRepo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to create PostgreSQL repository: %v", err)
	}
	mongoRepo, err := repository.NewMongoDBRepository(cfg.MongoURI, cfg.MongoDBName)
	if err != nil {
		log.Fatalf("Failed to create MongoDB repository: %v", err)
	}

	// Initialize services
	authService := services.NewAuthService(postgresRepo, cfg.JWTSecret, cfg.JWTExpiration, cfg.RefreshExpiration)
	profileService := services.NewProfileService(postgresRepo)
	aiService := services.NewAIService(postgresRepo, mongoRepo, cfg.OpenRouterKey)
	healthService := services.NewHealthService(postgresRepo)
	mediaService := services.NewMediaService(postgresRepo, mongoRepo)

	// Initialize handlers
	h := handlers.NewHandlers(authService, profileService, aiService, healthService, mediaService)

	// Setup router
	r := mux.NewRouter()

	// Middleware
	r.Use(middleware.LoggingMiddleware)

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Public routes
	r.HandleFunc("/health", h.HealthCheck).Methods("GET")
	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/refresh", h.RefreshToken).Methods("POST")

	// Authenticated routes
	authRouter := r.PathPrefix("/api").Subrouter()
	authRouter.Use(h.AuthMiddleware)
	{
		authRouter.HandleFunc("/profile", h.SaveProfile).Methods("POST")
		authRouter.HandleFunc("/profile", h.GetProfile).Methods("GET")
		authRouter.HandleFunc("/chat", h.Chat).Methods("POST")
		authRouter.HandleFunc("/chat/history", h.GetChatHistory).Methods("GET")
		authRouter.HandleFunc("/generate-plan", h.GeneratePlan).Methods("POST")
		authRouter.HandleFunc("/workout-plan", h.GetWorkoutPlan).Methods("GET")
		authRouter.HandleFunc("/regenerate-plan", h.RegenerateWorkoutPlan).Methods("POST")
		authRouter.HandleFunc("/complete-workout", h.CompleteWorkout).Methods("POST")
		authRouter.HandleFunc("/progress", h.GetUserProgress).Methods("GET")

		// Exercise media routes
		authRouter.HandleFunc("/exercise/{exercise_id}/media", h.GetExerciseMedia).Methods("GET")
		authRouter.HandleFunc("/exercise/media", h.SaveExerciseMedia).Methods("POST")
		authRouter.HandleFunc("/exercise/media/{media_id}", h.DeleteExerciseMedia).Methods("DELETE")
		authRouter.HandleFunc("/rating", h.GetRating).Methods("GET")
		authRouter.HandleFunc("/motivation", h.GetMotivationalMessage).Methods("GET")
	}

	// Start server
	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 300 * time.Second,
		IdleTimeout:  180 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
	log.Println("Server shutdown gracefully")
}

// runMigrations executes database migrations
func runMigrations(databaseURL string) error {
	// Use the migrations directory in the current working directory
	migrationsPath := "file://migrations"

	m, err := migrate.New(migrationsPath, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}
	defer m.Close()

	// Run all available migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Database migrations applied successfully")
	return nil
}
