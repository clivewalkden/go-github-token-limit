package main

import (
	"fmt"
	"go-github-token-limit/internal/utils"
	"net/http"
	"os"
	"time"

	"go-github-token-limit/internal/githubapi"
)

var version = "none provided"

func main() {
	fmt.Print("\033[H\033[2J") //clear screen
	println("")
	utils.InfoNotice(`GitHub Token Limit Checker`)
	utils.InfoNotice(fmt.Sprintf(`v%s`, version))
	println("")

	client := &http.Client{}
	token := githubapi.GetGithubTokenFromEnv()

	rateLimitResponse, err := githubapi.FetchRateLimit(client, token)
	if err != nil {
		utils.ErrorNotice(fmt.Sprintf("Error fetching rate limit: %v\n", err))
		os.Exit(3)
	}

	// Default limit is 60 if no token is provided
	if rateLimitResponse.Resources.Core.Limit == 60 {
		utils.CautionNotice("Please provide a GitHub token to in the environment variable " + githubapi.DefaultTokenEnv + ".")
		os.Exit(3)
	}

	core := rateLimitResponse.Resources.Core
	resetTime := core.Reset.Time
	//fmt.Printf("%+v\n", core)

	utils.SuccessNotice(fmt.Sprintf("Using token: %s", utils.ObscureToken(token)))

	if core.Remaining > 0 {
		utils.SuccessNotice(fmt.Sprintf("You have %d/%d requests left this hour", core.Remaining, core.Limit))
	} else {
		now := time.Now()
		durationUntilReset := resetTime.Sub(now).Minutes()
		utils.CautionNotice("You have no requests left.")
		utils.CautionNotice(fmt.Sprintf("The limit will reset in %.0f minutes at %s", durationUntilReset, resetTime))
	}
	println("")
}
