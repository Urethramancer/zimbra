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
	Users CmdUsers `command:"users" aliases:"u,usr" help:"User account queries and management."`
	Alias CmdAlias `command:"alias" aliases:"al" help:"Alias queries and management."`
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
