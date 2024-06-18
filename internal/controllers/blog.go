package controllers

import (
	"context"
	"errors"
	"github.com/sergeysergeevru/go-blog-service/internal/controllers/oapi"
	appErrors "github.com/sergeysergeevru/go-blog-service/internal/errors"
	"github.com/sergeysergeevru/go-blog-service/internal/models"
)

type BlogServiceProvider interface {
	CreatePost(ctx context.Context, post *models.Post) error
	UpdatePost(ctx context.Context, post *models.Post) error
	GetPostByID(ctx context.Context, id int) (*models.Post, error)
	DeletePost(ctx context.Context, id int) error
	GetPosts(ctx context.Context) ([]models.Post, error)
}

type PostController struct {
	service BlogServiceProvider
}

func NewPost(service BlogServiceProvider) *PostController {
	return &PostController{
		service: service,
	}
}

func (p PostController) GetPosts(ctx context.Context, request oapi.GetPostsRequestObject) (oapi.GetPostsResponseObject, error) {
	posts, err := p.service.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	var blogPosts oapi.GetPosts200JSONResponse
	for _, post := range posts {
		p := oapi.BlogPost{
			Author:  post.Author,
			Content: post.Content,
			Id:      &post.ID,
			Title:   post.Title,
		}
		blogPosts = append(blogPosts, p)
	}
	return blogPosts, nil
}

func (p PostController) PostPosts(ctx context.Context, request oapi.PostPostsRequestObject) (oapi.PostPostsResponseObject, error) {
	post := &models.Post{
		Title:   request.Body.Title,
		Content: request.Body.Content,
		Author:  request.Body.Author,
	}
	err := p.service.CreatePost(ctx, post)
	var validationError *appErrors.PostValidationError
	if errors.As(err, &validationError) {
		return oapi.PostPosts400JSONResponse{
			ErrorMessage: validationError.Error(),
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return oapi.PostPosts201JSONResponse{
		Author:  post.Author,
		Content: post.Content,
		Id:      &post.ID,
		Title:   post.Title,
	}, nil
}

func (p PostController) PutPostsId(ctx context.Context, request oapi.PutPostsIdRequestObject) (oapi.PutPostsIdResponseObject, error) {
	post := &models.Post{
		ID:      request.Id,
		Title:   request.Body.Title,
		Content: request.Body.Content,
		Author:  request.Body.Author,
	}
	err := p.service.UpdatePost(ctx, post)
	var validationError *appErrors.PostValidationError
	if errors.Is(err, appErrors.ErrPostNotFound) {
		return oapi.PutPostsId404Response{}, nil
	}
	if errors.As(err, &validationError) {
		return oapi.PutPostsId400JSONResponse{
			ErrorMessage: validationError.Error(),
		}, nil
	}
	if err != nil {
		return nil, err
	}
	return oapi.PutPostsId200JSONResponse{
		Author:  post.Author,
		Content: post.Content,
		Id:      &post.ID,
		Title:   post.Title,
	}, nil

}

func (p PostController) DeletePostsId(ctx context.Context, request oapi.DeletePostsIdRequestObject) (oapi.DeletePostsIdResponseObject, error) {
	err := p.service.DeletePost(ctx, request.Id)
	if errors.Is(err, appErrors.ErrPostNotFound) {
		return oapi.DeletePostsId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}
	return oapi.DeletePostsId204Response{}, nil
}

func (p PostController) GetPostsId(ctx context.Context, request oapi.GetPostsIdRequestObject) (oapi.GetPostsIdResponseObject, error) {
	post, err := p.service.GetPostByID(ctx, request.Id)
	if errors.Is(err, appErrors.ErrPostNotFound) {
		return oapi.GetPostsId404Response{}, nil
	}
	if err != nil {
		return nil, err
	}
	return oapi.GetPostsId200JSONResponse{
		Author:  post.Author,
		Content: post.Content,
		Id:      &post.ID,
		Title:   post.Title,
	}, nil
}
