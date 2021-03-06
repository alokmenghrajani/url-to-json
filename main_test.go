package main

import (
	"bou.ke/monkey"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	os.Args = []string{"git-to-json"}
	assert.Panics(t, func() { main() })

	os.Args = []string{"git-to-json", "/foobar"}
	assert.NotPanics(t, func() { main() })
}

func TestHttpsUrl(t *testing.T) {
	json := urlToJson("https://joe:hunter@example.com:443/my///pa/th/foo.html?a=1&b=2#magic")
	assert.Equal(t, "https", json.Scheme)
	assert.Equal(t, "joe", json.Username)
	assert.Equal(t, "hunter", json.Password)
	assert.Equal(t, "example.com", json.Host)
	assert.Equal(t, uint16(443), json.Port)
	assert.Equal(t, "my", json.Path[0])
	assert.Equal(t, "pa", json.Path[1])
	assert.Equal(t, "th", json.Path[2])
	assert.Equal(t, "foo.html", json.Path[3])
	assert.Equal(t, "1", json.Query["a"][0])
	assert.Equal(t, "2", json.Query["b"][0])
	assert.Equal(t, "magic", json.Fragment)
}

func TestGitUrl(t *testing.T) {
	json := urlToJson("org-12345678@github.com:alokmenghrajani/url-to-json.git")
	assert.Equal(t, "ssh", json.Scheme)
	assert.Equal(t, "org-12345678", json.Username)
	assert.Equal(t, "github.com", json.Host)
	assert.Equal(t, "alokmenghrajani", json.Path[0])
	assert.Equal(t, "url-to-json.git", json.Path[1])
}

func TestPanicOnError(t *testing.T) {
	assert.Panics(t, func() { panicOnError(errors.New("foobar")) })
	assert.NotPanics(t, func() { panicOnError(nil) })
}
