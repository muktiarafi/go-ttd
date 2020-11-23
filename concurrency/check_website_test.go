package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "anjay.com" {
		return false
	}
	return true
}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"anjay.com",
	}

	want := map[string]bool{
		"https://google.com":   true,
		"https://facebook.com": true,
		"anjay.com":            false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
}
