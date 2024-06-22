package githubapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	APIURL          = "https://api.github.com/rate_limit"
	AuthHeader      = "Authorization"
	DefaultTokenEnv = "GITHUB_TOKEN"
)

var TokenEnvNames = []string{
	"GITHUB_TOKEN",
	"GH_TOKEN",
	"GITHUB_ACCESS_TOKEN",
	"GH_ACCESS_TOKEN",
	"GITHUB_OAUTH_TOKEN",
	"GH_OAUTH_TOKEN",
	"GITHUB_PAT",
	"GH_PAT",
	"GITHUB_AUTH_TOKEN",
	"GH_AUTH_TOKEN",
	"GITHUB_API_TOKEN",
	"GH_API_TOKEN",
	"GITHUB_API_KEY",
	"GH_API_KEY",
	"GITHUB_PERSONAL_ACCESS_TOKEN",
	"GH_PERSONAL_ACCESS_TOKEN",
	"GITHUB_PERSONAL_TOKEN",
	"GH_PERSONAL_TOKEN",
	"GITHUB_PERSONAL_API_TOKEN",
	"GH_PERSONAL_API_TOKEN",
	"GITHUB_PERSONAL_API_KEY",
	"GH_PERSONAL_API_KEY",
	"GITHUB_APP_TOKEN",
	"GH_APP_TOKEN",
	"GITHUB_APP_KEY",
	"GH_APP_KEY",
}

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

func GetGithubTokenFromEnv() string {
	for _, envName := range TokenEnvNames {
		if token := os.Getenv(envName); token != "" {
			return token
		}
	}
	return ""
}
