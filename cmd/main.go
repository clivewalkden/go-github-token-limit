package main

import (
	"encoding/json"
	"fmt"
	"go-github-token-limit/internal/utils"
	"net/http"
	"os"
	"time"
)

const (
	apiUrl       = "https://api.github.com/rate_limit"
	authHeader   = "Authorization"
	tokenEnvName = "GITHUB_TOKEN"
)

type Rate struct {
	Limit     int       `json:"limit"`
	Remaining int       `json:"remaining"`
	Reset     Timestamp `json:"reset"`
}

type RateLimit struct {
	Core Rate `json:"core"`
}

type RateLimitResponse struct {
	Resources RateLimit `json:"resources"`
}

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(data []byte) error {
	// The data comes as a JSON number (UNIX timestamp)
	var unixTime int64
	if err := json.Unmarshal(data, &unixTime); err != nil {
		return err
	}
	*t = Timestamp(time.Unix(unixTime, 0))
	return nil
}

func main() {
	utils.InfoNotice(`                   GitHub Token Limit Checker                   `)
	utils.InfoNotice(`                             v1.0.0                             `)
	println("") // This is a placeholder for the main function

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		utils.ErrorNotice(` ERROR          `)
		utils.ErrorNotice(err.Error())
		os.Exit(3)
	}

	token := os.Getenv(tokenEnvName)
	if token != "" {
		req.Header.Set(authHeader, fmt.Sprintf("token %s", token))
	}

	resp, err := client.Do(req)
	if err != nil {
		utils.ErrorNotice(` ERROR          `)
		utils.ErrorNotice(err.Error())
		os.Exit(3)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		utils.ErrorNotice(` ERROR          `)
		utils.ErrorNotice(`Failed to get rate limit`)
		os.Exit(3)
	}

	var rateLimit RateLimitResponse
	if err := json.NewDecoder(resp.Body).Decode(&rateLimit); err != nil {
		utils.ErrorNotice(` ERROR          `)
		utils.ErrorNotice(err.Error())
		os.Exit(3)
	}

	core := rateLimit.Resources.Core
	resetTime := time.Time(core.Reset)
	//fmt.Printf("%+v\n", core)

	if core.Remaining > 0 {
		utils.InfoNotice(fmt.Sprintf("     You have %d/%d requests left this hour", core.Remaining, core.Limit))
	} else {
		now := time.Now()
		durationUntilReset := resetTime.Sub(now).Minutes()
		utils.InfoNotice(fmt.Sprintf("     You have no requests left. The limit will reset in %.0f at %s", durationUntilReset, resetTime))
	}
	println("")
}
