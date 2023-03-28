package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainPage(t *testing.T) {
	assert := assert.New(t)
	ah := NewRouter()

	ts := httptest.NewServer(ah)
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
}
