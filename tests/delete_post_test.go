package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlog_DeletePost(t *testing.T) {
	basicPost := `{"title": "test article", "content":"content description", "author": "Some author"}`
	post := createPostForTest(t, basicPost, 8080)

	tests := []struct {
		name               string
		postID             int
		expectedStatusCode int
		expectedResponse   *string
	}{
		{
			name:               "post not found",
			postID:             12321,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "post found",
			expectedStatusCode: http.StatusNoContent,
			postID:             post.ID,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8080/api/v1/posts/%d", test.postID), nil)
			assert.NoError(t, err)
			response, err := http.DefaultClient.Do(request)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			bodyBytes, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			if test.expectedResponse != nil {
				assert.JSONEq(t, *test.expectedResponse, string(bodyBytes))
			}
		})

	}
}
