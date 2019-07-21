package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/zimbra"
)

type CmdAlias struct {
	opt.DefaultHelp
	List CmdAliasList `command:"list" aliases:"ls" help:"List aliases for user accounts."`
}

func (cmd *CmdAlias) Run(in []string) error {
	return errors.New(opt.ErrorUsage)
}

type CmdAliasList struct {
	opt.DefaultHelp
	Email []string `placeholder:"EMAIL" help:"User account(s) to show aliases for."`
}

func (cmd *CmdAliasList) Run(in []string) error {
	if cmd.Help || len(cmd.Email) == 0 {
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
	for _, acc := range cmd.Email {
		list, err := zc.GetAliases(acc)
		if err != nil {
			return err
		}

		if len(list) > 0 {
			list = list[1:]
		}
		if len(list) == 0 {
			m("No aliases for %s", acc)
		} else {
			m("%s has %d aliases:", acc, len(list))
			for _, a := range list {
				m("\t%s", a)
			}
		}
	}
	return nil
}
