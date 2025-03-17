package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"personal-blog/internal/entity"
	"sync"
)

// ArticleRepository defines the interface for managing articles
type ArticleRepository interface {
	GetAll() ([]entity.Article, error)         // Retrieve all articles
	GetByID(id string) (entity.Article, error) // Retrieve an article by ID
	Save(article entity.Article) error         // Save a new or updated article
	Delete(id string) error                    // Delete an article by ID
}

// FileArticleRepository implements ArticleRepository using the filesystem
type FileArticleRepository struct {
	dirPath string // Directory where articles are stored
	mu      sync.Mutex
}

// NewFileArticleRepository creates a new instance of FileArticleRepository
func NewFileArticleRepository(dirPath string) *FileArticleRepository {
	return &FileArticleRepository{dirPath: dirPath}
}

// GetAll retrieves all articles from the filesystem
func (r *FileArticleRepository) GetAll() ([]entity.Article, error) {
	files, err := ioutil.ReadDir(r.dirPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var articles []entity.Article
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := ioutil.ReadFile(filepath.Join(r.dirPath, file.Name()))
			if err != nil {
				continue
			}
			var article entity.Article
			json.Unmarshal(data, &article)
			articles = append(articles, article)
		}
	}
	return articles, nil
}

// GetByID retrieves an article by its ID
func (r *FileArticleRepository) GetByID(id string) (entity.Article, error) {
	filePath := filepath.Join(r.dirPath, id+".json")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return entity.Article{}, fmt.Errorf("failed to read file: %w", err)
	}

	var article entity.Article
	json.Unmarshal(data, &article)
	return article, nil
}

// Save stores an article in the filesystem
func (r *FileArticleRepository) Save(article entity.Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join(r.dirPath, article.ID+".json")
	data, err := json.MarshalIndent(article, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

// Delete removes an article by its ID
func (r *FileArticleRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	filePath := filepath.Join(r.dirPath, id+".json")
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
