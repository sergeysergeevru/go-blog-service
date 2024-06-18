package app

import (
	"context"
	"errors"
	"fmt"
	errors2 "github.com/sergeysergeevru/go-blog-service/internal/errors"
	"github.com/sergeysergeevru/go-blog-service/internal/models"
	"github.com/sergeysergeevru/go-blog-service/internal/repositories"
)

type BlogService struct {
	postProvider PostRepository
}

//go:generate mockgen -source blog.go -destination blog_mock_test.go -package app
type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
	GetByID(ctx context.Context, id int) (*models.Post, error)
	Update(ctx context.Context, post *models.Post) error
	Delete(ctx context.Context, post *models.Post) error
	GetList(ctx context.Context) ([]models.Post, error)
}

func (b *BlogService) GetPosts(ctx context.Context) ([]models.Post, error) {
	return b.postProvider.GetList(ctx)
}

func (b *BlogService) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := b.postProvider.GetByID(ctx, post.ID)
	if errors.Is(err, repositories.ErrPostNotFound) {
		return errors2.ErrPostNotFound
	}
	if err != nil {
		return fmt.Errorf("can not get post by id: %w", err)
	}
	if err := models.ValidatePost(*post); err != nil {
		return fmt.Errorf("post content is not valid: %w", err)
	}
	return b.postProvider.Update(ctx, post)
}

func (b *BlogService) GetPostByID(ctx context.Context, id int) (*models.Post, error) {
	post, err := b.postProvider.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrPostNotFound) {
		return nil, errors2.ErrPostNotFound
	}
	return post, nil
}

func NewBlogService(repository PostRepository) *BlogService {
	return &BlogService{postProvider: repository}
}

func (b *BlogService) CreatePost(ctx context.Context, post *models.Post) error {
	if err := models.ValidatePost(*post); err != nil {
		return fmt.Errorf("post content is not valid: %w", err)
	}
	return b.postProvider.Create(ctx, post)
}

func (b *BlogService) DeletePost(ctx context.Context, id int) error {
	post, err := b.postProvider.GetByID(ctx, id)
	if errors.Is(err, repositories.ErrPostNotFound) {
		return errors2.ErrPostNotFound
	}
	if err != nil {
		return fmt.Errorf("DeletePost: can not get post by id: %w", err)
	}
	err = b.postProvider.Delete(ctx, post)
	if err != nil {
		return fmt.Errorf("can not delete a post: %w", err)
	}
	return nil
}
