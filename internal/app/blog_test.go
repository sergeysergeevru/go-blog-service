package app

import (
	"context"
	"errors"
	"fmt"
	errors2 "github.com/sergeysergeevru/go-blog-service/internal/errors"
	"github.com/sergeysergeevru/go-blog-service/internal/models"
	"github.com/sergeysergeevru/go-blog-service/internal/repositories"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestBlogService_UpdatePost(t *testing.T) {
	var (
		ErrGetPost    = fmt.Errorf("get post error")
		ErrUpdatePost = fmt.Errorf("update post error")
	)
	tests := []struct {
		name       string
		post       *models.Post
		setUpMock  func(postProvider *MockPostRepository)
		checkError func(t *testing.T, err error)
	}{
		{
			name: "post with same ID does not exist",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(nil, repositories.ErrPostNotFound)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, errors2.ErrPostNotFound)
			},
		},
		{
			name: "internal error for getting post",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(nil, ErrGetPost)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, ErrGetPost)
			},
		},
		{
			name: "post exist, all fields empty",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(&models.Post{}, nil)
			},
			checkError: func(t *testing.T, err error) {
				var validationError *errors2.PostValidationError
				require.True(t, errors.As(err, &validationError))
			},
		},
		{
			name: "error on save post",
			post: &models.Post{
				ID:      1,
				Title:   "title",
				Author:  "tset",
				Content: "test",
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(&models.Post{}, nil)
				postProvider.EXPECT().Update(gomock.Any(), postUpdate).Return(ErrUpdatePost)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, ErrUpdatePost)
			},
		},
		{
			name: "successfully update",
			post: &models.Post{
				ID:      1,
				Title:   "title",
				Author:  "tset",
				Content: "test",
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(&models.Post{}, nil)
				postProvider.EXPECT().Update(gomock.Any(), postUpdate).Return(nil)
			},
			checkError: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			postProvider := NewMockPostRepository(ctrl)
			tt.setUpMock(postProvider)
			b := NewBlogService(postProvider)
			err := b.UpdatePost(context.Background(), tt.post)
			tt.checkError(t, err)
		})
	}
}
func TestBlogService_DeletePost(t *testing.T) {
	var (
		ErrGetPost    = fmt.Errorf("get post error")
		ErrDeletePost = fmt.Errorf("delete post error")
	)
	tests := []struct {
		name       string
		post       *models.Post
		setUpMock  func(postProvider *MockPostRepository)
		checkError func(t *testing.T, err error)
	}{
		{
			name: "post with same ID does not exist",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(nil, repositories.ErrPostNotFound)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, errors2.ErrPostNotFound)
			},
		},
		{
			name: "internal error for getting post",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(nil, ErrGetPost)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, ErrGetPost)
			},
		},
		{
			name: "error on delete post",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(postUpdate, nil)
				postProvider.EXPECT().Delete(gomock.Any(), postUpdate).Return(ErrDeletePost)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, ErrDeletePost)
			},
		},
		{
			name: "successfully deleted",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().
					GetByID(gomock.Any(), 1).Return(postUpdate, nil)
				postProvider.EXPECT().Delete(gomock.Any(), postUpdate).Return(nil)
			},
			checkError: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			postProvider := NewMockPostRepository(ctrl)
			tt.setUpMock(postProvider)
			b := NewBlogService(postProvider)
			err := b.DeletePost(context.Background(), tt.post.ID)
			tt.checkError(t, err)
		})
	}
}

func TestBlogService_CreatePost(t *testing.T) {
	var (
		ErrCreatePost = fmt.Errorf("create post error")
	)
	tests := []struct {
		name       string
		post       *models.Post
		setUpMock  func(postProvider *MockPostRepository)
		checkError func(t *testing.T, err error)
	}{
		{
			name: "validation error",
			post: &models.Post{
				ID: 1,
			},
			setUpMock: func(postProvider *MockPostRepository) {},
			checkError: func(t *testing.T, err error) {
				var validationError *errors2.PostValidationError
				require.True(t, errors.As(err, &validationError))
			},
		},
		{
			name: "error on save post",
			post: &models.Post{
				ID:      1,
				Title:   "title",
				Author:  "tset",
				Content: "test",
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().Create(gomock.Any(), postUpdate).Return(ErrCreatePost)
			},
			checkError: func(t *testing.T, err error) {
				require.ErrorIs(t, err, ErrCreatePost)
			},
		},
		{
			name: "successfully update",
			post: &models.Post{
				ID:      1,
				Title:   "title",
				Author:  "tset",
				Content: "test",
			},
			setUpMock: func(postProvider *MockPostRepository) {
				postUpdate := &models.Post{
					ID:      1,
					Title:   "title",
					Author:  "tset",
					Content: "test",
				}
				postProvider.EXPECT().Create(gomock.Any(), postUpdate).Return(nil)
			},
			checkError: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			postProvider := NewMockPostRepository(ctrl)
			tt.setUpMock(postProvider)
			b := NewBlogService(postProvider)
			err := b.CreatePost(context.Background(), tt.post)
			tt.checkError(t, err)
		})
	}
}
