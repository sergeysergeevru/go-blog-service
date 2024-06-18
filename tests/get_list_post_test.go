package tests

import (
	"fmt"
	"github.com/sergeysergeevru/go-blog-service/internal/controllers"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlog_GetPostList(t *testing.T) {
	//run dedicated server since the main one is flaky
	root := controllers.CreateHandler(basicCfg)
	go func() {
		root.Run(":8081")
	}()
	basicPost := `{"title": "test article", "content":"content description", "author": "Some author"}`
	post := createPostForTest(t, basicPost, 8081)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedResponse   *string
	}{
		{
			name:               "post found",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   getStringPointer(fmt.Sprintf(`[{"id":%d,"title": "test article", "content":"content description", "author": "Some author"}]`, post.ID)),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := http.NewRequest(http.MethodGet, "http://localhost:8081/api/v1/posts", nil)
			assert.NoError(t, err)
			response, err := http.DefaultClient.Do(request)
			assert.NoError(t, err)
			assert.Equal(t, test.expectedStatusCode, response.StatusCode)
			bodyBytes, err := io.ReadAll(response.Body)
			t.Log(string(bodyBytes))
			assert.NoError(t, err)
			if test.expectedResponse != nil {
				assert.JSONEq(t, *test.expectedResponse, string(bodyBytes))
			}
		})

	}
}
