package main

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "test-key"
	val := []byte("test-data")

	cache.Add(key, val)

	result, ok := cache.Get(key)
	if !ok {
		t.Error("expected to find key")
	}

	if string(result) != string(val) {
		t.Error("expected values to match")
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
