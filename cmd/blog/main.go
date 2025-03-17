package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"personal-blog/internal/auth"
	"personal-blog/internal/handler"
	"personal-blog/internal/repository"
	"personal-blog/internal/usecase"
)

func main() {
	// Ensure the articles directory exists
	articlesDir := "./articles"
	if err := os.MkdirAll(articlesDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create articles directory: %v", err)
	}

	// Initialize repository
	articleRepo := repository.NewFileArticleRepository(articlesDir)

	// Initialize use case
	articleUsecase := usecase.NewArticleUsecase(articleRepo)

	// Initialize HTTP handler
	httpHandler := handler.NewHTTPHandler(articleUsecase)

	// Define routes
	http.HandleFunc("/", httpHandler.HomeHandler)                                                    // Home page
	http.HandleFunc("/article/", httpHandler.ArticleHandler)                                         // Article page
	http.HandleFunc("/admin/dashboard", auth.AdminMiddleware(httpHandler.AdminDashboardHandler))     // Admin dashboard
	http.HandleFunc("/admin/add-article", auth.AdminMiddleware(httpHandler.AddArticleHandler))       // Add article
	http.HandleFunc("/admin/edit-article", auth.AdminMiddleware(httpHandler.EditArticleHandler))     // Edit article
	http.HandleFunc("/admin/delete-article", auth.AdminMiddleware(httpHandler.DeleteArticleHandler)) // Delete article
	http.HandleFunc("/admin/login", auth.LoginHandler)                                               // Admin login
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))     // Serve static files

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
