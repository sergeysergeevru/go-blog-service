package tests

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlog_CreatePost(t *testing.T) {
	tests := []struct {
		name               string
		body               string
		expectedStatusCode int
		expectedResponse   *string
	}{
		{
			name:               "non json body",
			body:               "",
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "empty fields of the request",
			body:               "{}",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message": "title field is empty"}`),
		},
		{
			name:               "title in place, the rest fields are empty",
			body:               `{"title": "test article"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message": "content field is empty"}`),
		},
		{
			name:               "title, content in place, the rest fields are empty",
			body:               `{"title": "test article", "content":"content description"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   getStringPointer(`{"error_message": "author field is empty"}`),
		},
		{
			name:               "success on creation",
			body:               `{"title": "test article", "content":"content description", "author": "Some author"}`,
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   getStringPointer(`{"id":1,"title": "test article", "content":"content description", "author": "Some author"}`),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/api/v1/posts", strings.NewReader(test.body))
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
