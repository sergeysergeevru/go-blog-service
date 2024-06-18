package tests

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"strings"
	"testing"
)

func getStringPointer(str string) *string {
	return &str
}

type TestPost struct {
	ID int `json:"id"`
}

func createPostForTest(t *testing.T, basicPost string, port int) TestPost {
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%d/api/v1/posts", port), strings.NewReader(basicPost))
	assert.NoError(t, err)
	response, err := http.DefaultClient.Do(request)
	require.Equal(t, http.StatusCreated, response.StatusCode)
	require.NoError(t, err)
	basicPostResponse, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	post := TestPost{}
	err = json.Unmarshal(basicPostResponse, &post)
	require.NoError(t, err)
	return post
}
