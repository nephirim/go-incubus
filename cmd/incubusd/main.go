package main

import (
	"github.com/nephirim/go-incubus/app/params"
	"os"

	"github.com/nephirim/go-incubus/app"
	"github.com/nephirim/go-incubus/cmd/incubusd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	params.SetAddressPrefixes()

	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
