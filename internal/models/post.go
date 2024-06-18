package models

import "github.com/sergeysergeevru/go-blog-service/internal/errors"

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func ValidatePost(post Post) error {
	if post.Title == "" {
		return errors.NewPostValidationError("title field is empty")
	}
	if post.Content == "" {
		return errors.NewPostValidationError("content field is empty")
	}
	if post.Author == "" {
		return errors.NewPostValidationError("author field is empty")
	}
	return nil
}
