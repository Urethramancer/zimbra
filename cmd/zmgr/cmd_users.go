package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/zimbra"
)

type CmdUsers struct {
	opt.DefaultHelp
	List CmdUsersList `command:"list" aliases:"ls" help:"List user accounts."`
}

func (cmd *CmdUsers) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

type CmdUsersList struct {
	opt.DefaultHelp
	Domain string `placeholder:"DOMAIN" help:"Optional domain to limit search."`
}

func (cmd *CmdUsersList) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	cfg := loadConfig()
	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cfg.Admin, cfg.Password)
	m := log.Default.Msg
	if err != nil {
		m("Login failed. (%s)", err.Error())
		return nil
	}

	defer zc.Close()
	var list []string
	if cmd.Domain == "" {
		list, err = zc.GetUsers()
	} else {
		list, err = zc.GetUsersInDomain(cmd.Domain)
	}
	if err != nil {
		return err
	}

	for _, x := range list {
		m("%s", x)
	}
	return nil
}
