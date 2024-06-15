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
		utils.SuccessNotice(fmt.Sprintf("You have %d/%d requests left this hour", core.Remaining, core.Limit))
	} else {
		now := time.Now()
		durationUntilReset := resetTime.Sub(now).Minutes()
		utils.CautionNotice("You have no requests left.")
		utils.CautionNotice(fmt.Sprintf("The limit will reset in %.0f minutes at %s", durationUntilReset, resetTime))
	}
	println("")
}
