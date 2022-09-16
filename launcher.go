package main

import (
	"github.com/fatih/color"
	"kamigen/2fucked2pay/src/bots"
	"kamigen/2fucked2pay/src/utils"
)

const (
	cid = "88650e7e-efee-4857-b9a9-cf580a00ef43"
)

func main() {
	color.Cyan("Welcome to 2fucked2pay launcher")
	color.Cyan("You have to login to your Microsoft account using the link auth")
	t, err := utils.GetMCcredentials("./cache", cid)
	if err != nil {
		color.Red("Error: %v", err)
		return
	}
	color.Green("Successfully logged in")
	bots.Join(&t)
}
