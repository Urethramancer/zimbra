package main

import (
	"os"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

var Options struct {
	opt.DefaultHelp
	Setup CmdSetup `command:"setup" help:"Set up basic settings."`
	Auth  CmdAuth  `command:"authenticate" aliases:"auth,login" help:"Authenticate a user to test that it works/exists."`
}

func main() {
	a := opt.Parse(&Options)
	if Options.Help || len(os.Args) < 2 {
		a.Usage()
		return
	}

	err := a.RunCommand(false)
	if err != nil {
		log.Default.Msg("Error running: %s", err.Error())
		os.Exit(2)
	}
}
