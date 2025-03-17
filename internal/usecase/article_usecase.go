package usecase

import (
	"errors"
	"strconv"
	"time"

	"personal-blog/internal/entity"
	"personal-blog/internal/repository"
)

// ArticleUsecase contains the business logic for managing articles
type ArticleUsecase struct {
	repo repository.ArticleRepository
}

// NewArticleUsecase creates a new instance of ArticleUsecase
func NewArticleUsecase(repo repository.ArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

// GetAllArticles retrieves all articles
func (uc *ArticleUsecase) GetAllArticles() ([]entity.Article, error) {
	return uc.repo.GetAll()
}

// GetArticleByID retrieves an article by its ID
func (uc *ArticleUsecase) GetArticleByID(id string) (entity.Article, error) {
	return uc.repo.GetByID(id)
}

// AddArticle adds a new article
func (uc *ArticleUsecase) AddArticle(title, content string) (string, error) {
	if title == "" || content == "" {
		return "", errors.New("title and content cannot be empty")
	}

	id := strconv.FormatInt(time.Now().UnixNano(), 10)
	article := entity.Article{
		ID:            id,
		Title:         title,
		Content:       content,
		PublishedDate: time.Now(),
	}

	err := uc.repo.Save(article)
	if err != nil {
		return "", err
	}
	return id, nil
}

// UpdateArticle updates an existing article
func (uc *ArticleUsecase) UpdateArticle(id, title, content string) error {
	article, err := uc.repo.GetByID(id)
	if err != nil {
		return err
	}

	article.Title = title
	article.Content = content
	return uc.repo.Save(article)
}

// DeleteArticle deletes an article by its ID
func (uc *ArticleUsecase) DeleteArticle(id string) error {
	return uc.repo.Delete(id)
}
