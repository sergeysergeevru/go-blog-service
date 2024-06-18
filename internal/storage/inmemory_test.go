package storage

import (
	"context"
	"fmt"
	"github.com/sergeysergeevru/go-blog-service/internal/models"
	"github.com/sergeysergeevru/go-blog-service/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

func TestCreateUpdateRaceCondition(t *testing.T) {
	postStorage := NewPostInMemory()
	wg := sync.WaitGroup{}
	ctx := context.Background()
	totalIterations := 1000
	for i := 0; i < totalIterations; i++ {
		wg.Add(1)
		go func(createIndex int) {
			postCreate := &models.Post{}
			postStorage.Create(ctx, postCreate)
			for j := 0; j < totalIterations; j++ {
				wg.Add(1)
				go func(createIndex int, updateIndex int) {
					postUpdate := &models.Post{
						ID:    postCreate.ID,
						Title: fmt.Sprintf("post %d update %d", createIndex, updateIndex),
					}
					postStorage.Update(ctx, postUpdate)
					wg.Done()
				}(createIndex, j)

			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	postList, err := postStorage.GetList(ctx)
	if err != nil {
		t.Fatalf("error received: %s", err)
	}
	if len(postList) != totalIterations {
		t.Fatalf("recevied %d posts", len(postList))
	}
}

func TestPosts_AllCommands(t *testing.T) {
	ctx := context.Background()
	storage := NewPostInMemory()
	post := models.Post{
		Title:   "title",
		Content: "content",
		Author:  "author",
	}
	require.Nil(t, storage.Create(ctx, &post))

	storedPost, err := storage.GetByID(ctx, post.ID)
	require.Nil(t, err)

	post.ID = storedPost.ID
	assert.Equal(t, &post, storedPost)

	post.Title = "new title"
	require.Nil(t, storage.Update(ctx, &post))

	storedPost, err = storage.GetByID(ctx, post.ID)
	require.Nil(t, err)
	require.Equal(t, &post, storedPost)
	posts, err := storage.GetList(ctx)
	require.Nil(t, err)
	require.Equal(t, []models.Post{*storedPost}, posts)

	require.NoError(t, storage.Delete(ctx, storedPost))

	storedPost, err = storage.GetByID(ctx, post.ID)
	require.Nil(t, storedPost)
	require.ErrorIs(t, err, repositories.ErrPostNotFound)
}
