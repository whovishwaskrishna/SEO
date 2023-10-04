package urlstorage_test

import (
	"github.com/stjudewashere/seonaut/internal/urlstorage"
	"testing"
)

func TestURLStorage(t *testing.T) {
	s := urlstorage.New()
	url := "http://example.com"

	// Test the Seen method
	if s.Seen(url) {
		t.Errorf("Expected %s to not be seen", url)
	}

	// Test the Add and Seen methods
	s.Add(url)
	if !s.Seen(url) {
		t.Errorf("Expected %s to be seen", url)
	}

	// Test the Iterate method
	seen := make(map[string]bool)
	s.Iterate(func(u string) {
		seen[u] = true
	})
	if !seen[url] {
		t.Errorf("Expected %s to be in the seen map", url)
	}
}
