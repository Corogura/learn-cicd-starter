package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	want := "TestAPIKey"
	header := http.Header{}
	header.Set("Authorization", "ApiKey"+want)
	got, _ := GetAPIKey(header)
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAPIKeyError(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "NotAPIKey")
	_, err := GetAPIKey(header)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
