package handler

import (
	"html/template"
	"net/http"
	"strings"

	"personal-blog/internal/usecase"
)

// HTTPHandler manages HTTP requests
type HTTPHandler struct {
	usecase *usecase.ArticleUsecase
}

// NewHTTPHandler creates a new instance of HTTPHandler
func NewHTTPHandler(usecase *usecase.ArticleUsecase) *HTTPHandler {
	return &HTTPHandler{usecase: usecase}
}

// HomeHandler displays the list of articles
func (h *HTTPHandler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := h.usecase.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to load articles", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	tmpl.Execute(w, articles)
}

// ArticleHandler displays a single article
func (h *HTTPHandler) ArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/article/")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	article, err := h.usecase.GetArticleByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/article.html"))
	tmpl.Execute(w, article)
}

// AdminDashboardHandler displays the admin dashboard
func (h *HTTPHandler) AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := h.usecase.GetAllArticles()
	if err != nil {
		http.Error(w, "Failed to load articles", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/admin_dashboard.html"))
	tmpl.Execute(w, articles)
}

// AddArticleHandler handles adding a new article
func (h *HTTPHandler) AddArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := h.usecase.AddArticle(title, content)
		if err != nil {
			http.Error(w, "Failed to add article", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/add_article.html"))
	tmpl.Execute(w, nil)
}

// EditArticleHandler handles editing an existing article
func (h *HTTPHandler) EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		title := r.FormValue("title")
		content := r.FormValue("content")

		err := h.usecase.UpdateArticle(id, title, content)
		if err != nil {
			http.Error(w, "Failed to update article", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/admin/edit-article/")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	article, err := h.usecase.GetArticleByID(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/edit_article.html"))
	tmpl.Execute(w, article)
}

// DeleteArticleHandler handles deleting an article
func (h *HTTPHandler) DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/admin/delete-article/")
	if id == "" {
		http.NotFound(w, r)
		return
	}

	err := h.usecase.DeleteArticle(id)
	if err != nil {
		http.Error(w, "Failed to delete article", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
}
