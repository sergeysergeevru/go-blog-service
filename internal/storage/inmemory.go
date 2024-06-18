package storage

import (
	"context"
	"github.com/sergeysergeevru/go-blog-service/internal/models"
	"github.com/sergeysergeevru/go-blog-service/internal/repositories"
	"sync"
)

type Posts struct {
	sync.RWMutex
	data    map[int]models.Post
	counter int
}

func NewPostInMemory() *Posts {
	return &Posts{
		data: make(map[int]models.Post),
	}
}

func (p *Posts) Create(ctx context.Context, post *models.Post) error {
	p.Lock()
	defer p.Unlock()
	p.counter++
	post.ID = p.counter
	p.data[p.counter] = *post
	return nil
}

func (p *Posts) GetByID(ctx context.Context, id int) (*models.Post, error) {
	p.RLock()
	defer p.RUnlock()
	post, exist := p.data[id]
	if !exist {
		return nil, repositories.ErrPostNotFound
	}
	return &post, nil
}

func (p *Posts) Update(ctx context.Context, post *models.Post) error {
	p.Lock()
	defer p.Unlock()
	p.data[post.ID] = *post
	return nil
}

func (p *Posts) GetList(ctx context.Context) ([]models.Post, error) {
	p.RLock()
	defer p.RUnlock()
	posts := make([]models.Post, len(p.data))
	i := 0
	for k := range p.data {
		posts[i] = p.data[k]
		i++
	}
	return posts, nil
}

func (p *Posts) Delete(ctx context.Context, post *models.Post) error {
	p.Lock()
	defer p.Unlock()
	delete(p.data, post.ID)
	return nil
}
