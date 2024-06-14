package main

import (
	"encoding/json"
	"go-github-token-limit/internal/githubapi"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFetchRateLimit(t *testing.T) {
	// Mock server to simulate GitHub API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := githubapi.RateLimitResponse{
			Resources: githubapi.RateLimit{
				Core: githubapi.Rate{
					Limit:     5000,
					Remaining: 4999,
					Reset:     githubapi.Timestamp{Time: time.Now().Add(30 * time.Minute)},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	originalAPIURL := githubapi.APIURL
	githubapi.APIURL = mockServer.URL
	defer func() { githubapi.APIURL = originalAPIURL }()

	client := &http.Client{}
	token := ""

	rateLimitResponse, err := githubapi.FetchRateLimit(client, token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	core := rateLimitResponse.Resources.Core
	if core.Limit != 5000 {
		t.Errorf("Expected limit to be 5000, got %d", core.Limit)
	}
	if core.Remaining != 4999 {
		t.Errorf("Expected remaining to be 4999, got %d", core.Remaining)
	}
	if core.Reset.Time.Before(time.Now()) {
		t.Errorf("Expected reset time to be in the future, got %s", core.Reset.Time)
	}
}

func TestFetchRateLimitReached(t *testing.T) {
	// Mock server to simulate GitHub API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := githubapi.RateLimitResponse{
			Resources: githubapi.RateLimit{
				Core: githubapi.Rate{
					Limit:     5000,
					Remaining: 0,
					Reset:     githubapi.Timestamp{Time: time.Now().Add(30 * time.Minute)},
				},
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	originalAPIURL := githubapi.APIURL
	githubapi.APIURL = mockServer.URL
	defer func() { githubapi.APIURL = originalAPIURL }()

	client := &http.Client{}
	token := ""

	rateLimitResponse, err := githubapi.FetchRateLimit(client, token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	core := rateLimitResponse.Resources.Core
	if core.Limit != 5000 {
		t.Errorf("Expected limit to be 5000, got %d", core.Limit)
	}
	if core.Remaining != 0 {
		t.Errorf("Expected remaining to be 0, got %d", core.Remaining)
	}
	if core.Reset.Time.Before(time.Now()) {
		t.Errorf("Expected reset time to be in the future, got %s", core.Reset.Time)
	}
}
