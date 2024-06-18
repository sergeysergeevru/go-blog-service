package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlog_UpdatePost(t *testing.T) {
	basicPost := `{"title": "test article", "content":"content description", "author": "Some author"}`
	post := createPostForTest(t, basicPost, 8080)

	tests := []struct {
		name               string
		postID             int
		body               string
		expectedStatusCode int
		expectedResponse   *string
	}{
		{
			name:               "non json body",
			postID:             post.ID,
			body:               "",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "empty request object",
			postID:             post.ID,
			body:               "{}",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message":"title field is empty"}`),
		},
		{
			name:               "title in place, the rest fields are empty",
			postID:             post.ID,
			body:               `{"title": "update tests title"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message": "content field is empty"}`),
		},
		{
			name:               "title, content in place, the rest fields are empty",
			postID:             post.ID,
			body:               `{"title": "update test article", "content":"update content description"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message": "author field is empty"}`),
		},
		{
			name:               "success on update",
			postID:             post.ID,
			body:               `{"title": "update test article", "content":"update content description", "author": "update Some author"}`,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   getStringPointer(fmt.Sprintf(`{"id":%d,"title": "update test article", "content":"update content description", "author": "update Some author"}`, post.ID)),
		},
		{
			name:               "unknown post id",
			postID:             123213,
			body:               `{"title": "update test article", "content":"update content description", "author": "update Some author"}`,
			expectedStatusCode: http.StatusNotFound,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8080/api/v1/posts/%d", test.postID), strings.NewReader(test.body))
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
