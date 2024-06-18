package models

import (
	errors2 "github.com/sergeysergeevru/go-blog-service/internal/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidatePost(t *testing.T) {

	tests := []struct {
		name    string
		post    Post
		wantErr bool
	}{
		{
			name:    "all fields are empty",
			post:    Post{},
			wantErr: true,
		},
		{
			name: "title filled",
			post: Post{
				Title: "test",
			},
			wantErr: true,
		},
		{
			name: "title filled, author filled",
			post: Post{
				Title:  "test",
				Author: "test",
			},
			wantErr: true,
		},
		{
			name: "title filled, author filled, content filled",
			post: Post{
				Title:   "test",
				Author:  "test",
				Content: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidatePost(tt.post); (err != nil) != tt.wantErr {
				require.ErrorAs(t, err, errors2.PostValidationError{})
			}
		})
	}
}
