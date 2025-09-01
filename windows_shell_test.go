package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsShell = []struct {
	name    string
	err     bool
	browser Browser
}{
	// the table itself
	// {"default", false, ""},
	// {"chrome", false, Chrome},
	// {"firefox", false, Firefox},
	{"edge", false, Edge},
	{"yandex", false, Yandex},
}

func TestShell(t *testing.T) {
	// The execution loop
	// Capture tt for safety, use NoError, and put expected before actual in Equal
	for _, tt := range testsShell {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			url := `http://ya.ru`
			err := StartBrowser(tt.browser, url)
			if tt.err {
				assert.NotNil(t, err, "ожидаем ошибку")
			} else {
				// ожидаем отсутствие ошибки
				assert.NoError(t, err)
			}
		})
	}
}

func TestOpenHttpLinkInShell(t *testing.T) {
	url := `file://ya.ru`
	err := OpenHttpLinkInShell(url)
	assert.NoError(t, err, "ожидаем ошибку")
	url = `ya.ru`
	err = OpenHttpLinkInShell(url)
	assert.NoError(t, err, "ожидаем ошибку")
	url = `http://ya.ru`
	err = OpenHttpLinkInShell(url)
	assert.NoError(t, err, "ожидаем ошибку")
	url = `https://ya.ru`
	err = OpenHttpLinkInShell(url)
	assert.NotNil(t, err, "ожидаем ошибку")
	url = `//ya.ru`
	err = OpenHttpLinkInShell(url)
	assert.NoError(t, err, "ожидаем ошибку")
	url = `f//ya.ru`
	err = OpenHttpLinkInShell(url)
	assert.NoError(t, err, "ожидаем ошибку")
}
