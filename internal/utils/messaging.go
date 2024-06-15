package utils

import (
	"github.com/fatih/color"
)

func CautionNotice(message string) {
	infoMsg := color.New(color.FgHiYellow).Add(color.Bold).PrintlnFunc()
	infoMsg(CenterString(message, 80))
}

func InfoNotice(message string) {
	infoMsg := color.New(color.FgBlue).Add(color.Bold).PrintlnFunc()
	infoMsg(CenterString(message, 80))
}

func ErrorNotice(message string) {
	errorMsg := color.New(color.FgHiRed).Add(color.Bold).PrintlnFunc()
	errorMsg(CenterString(message, 80))
}

func SuccessNotice(message string) {
	successMsg := color.New(color.FgHiGreen).Add(color.Bold).PrintlnFunc()
	successMsg(CenterString(message, 80))
}
