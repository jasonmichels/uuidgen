package main

import (
	"regexp"
	"testing"
	"time"
)

func TestNewUUIDv4(t *testing.T) {
	uuid, err := NewUUIDv4()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(uuid) {
		t.Errorf("Invalid UUID v4: %s", uuid)
	}
}

func BenchmarkNewUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewUUIDv4()
	}
}

func TestUnixTime(t *testing.T) {
	unix := UnixTime()
	if unix <= 0 {
		t.Errorf("Unix timestamp should be a positive number")
	}

	// You can also compare it to a timestamp close to the time of the test
	expected := time.Now().Unix()
	if unix > expected+1 || unix < expected-1 {
		t.Errorf("Unix timestamp is too far off from the expected time")
	}
}

func BenchmarkUnixTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UnixTime()
	}
}
