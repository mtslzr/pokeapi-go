package pokeapi

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	scheme := "http"
	host := "test.io"

	err := SetAPIPath(fmt.Sprintf("%s://%s", scheme, host))

	assert.Nil(t, err)

	assert.Equal(t, host, apiHost)
	assert.Equal(t, scheme, apiScheme)
}

var dataTestSetAPIPath = []struct {
	Given        string
	ExpectHost   string
	ExpectScheme string
	ShouldErr    bool
}{
	{"http://foobar.com", "foobar.com", "http", false},
	{"https://zip.co", "zip.co", "https", false},
	{"whatisthis.com/", "whatisthis.com", "", true},
	{(&url.URL{Scheme: defaultAPIScheme, Host: defaultAPIHost}).String(), defaultAPIHost, defaultAPIScheme, false},
}

func TestAPIPath(t *testing.T) {
	for _, tt := range dataTestSetAPIPath {
		t.Run(tt.Given, func(t *testing.T) {
			err := SetAPIPath(tt.Given)

			assert.Equal(t, tt.ShouldErr, err != nil)
			if err != nil {
				return
			}

			assert.Equal(t, tt.ExpectHost, apiHost)
			assert.Equal(t, tt.ExpectScheme, apiScheme)
		})
	}
}
