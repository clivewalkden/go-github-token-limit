package main

import (
	"fmt"
	"go-github-token-limit/internal/utils"
	"net/http"
	"os"
	"time"

	"go-github-token-limit/internal/githubapi"
)

func main() {
	utils.InfoNotice(`                   GitHub Token Limit Checker                   `)
	utils.InfoNotice(`                             v1.0.0                             `)
	println("") // This is a placeholder for the main function

	client := &http.Client{}
	token := os.Getenv(githubapi.TokenEnvName)

	rateLimitResponse, err := githubapi.FetchRateLimit(client, token)
	if err != nil {
		utils.ErrorNotice(fmt.Sprintf("Error fetching rate limit: %v\n", err))
		os.Exit(3)
	}

	core := rateLimitResponse.Resources.Core
	resetTime := core.Reset.Time
	//fmt.Printf("%+v\n", core)

	if core.Remaining > 0 {
		utils.InfoNotice(fmt.Sprintf("     You have %d/%d requests left this hour", core.Remaining, core.Limit))
	} else {
		now := time.Now()
		durationUntilReset := resetTime.Sub(now).Minutes()
		utils.InfoNotice(fmt.Sprintf("     You have no requests left. The limit will reset in %.0f minute at %s", durationUntilReset, resetTime))
	}
	println("")
}
