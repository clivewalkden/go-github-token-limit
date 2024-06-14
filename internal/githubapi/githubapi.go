package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var (
	APIURL       = "https://api.github.com/rate_limit"
	AuthHeader   = "Authorization"
	TokenEnvName = "GITHUB_TOKEN"
)

type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     Timestamp `json:"reset"`
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var timestamp interface{}
	if err := json.Unmarshal(data, &timestamp); err != nil {
		return err
	}

	switch v := timestamp.(type) {
	case string:
		// Handle string representation of Unix timestamp
		parsedTime, err := time.Parse(time.RFC3339Nano, v)
		if err != nil {
			return err
		}
		t.Time = parsedTime
	case float64:
		// Handle numeric representation of Unix timestamp
		t.Time = time.Unix(int64(v), 0)
	default:
		return fmt.Errorf("unexpected type for timestamp: %T", v)
	}

	return nil
}

type RateLimit struct {
	Core Rate `json:"core"`
}

type RateLimitResponse struct {
	Resources RateLimit `json:"resources"`
}

func FetchRateLimit(client *http.Client, token string) (RateLimitResponse, error) {
	req, err := http.NewRequest("GET", APIURL, nil)
	if err != nil {
		return RateLimitResponse{}, err
	}

	if token != "" {
		req.Header.Set(AuthHeader, fmt.Sprintf("token %s", token))
	}

	resp, err := client.Do(req)
	if err != nil {
		return RateLimitResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RateLimitResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rateLimitResponse RateLimitResponse
	if err := json.NewDecoder(resp.Body).Decode(&rateLimitResponse); err != nil {
		return RateLimitResponse{}, err
	}

	return rateLimitResponse, nil
}
