package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/zimbra"
)

type CmdAuth struct {
	opt.DefaultHelp
	Email    string `placeholder:"EMAIL" help:"E-mail account to test logging in as. Leave blank to test admin login."`
	Password string `placeholder:"PASSWORD" help:"Password for the account."`
}

func (cmd *CmdAuth) Run(in []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	if cmd.Email == "" && cmd.Password == "" {
		cfg := loadConfig()
		cmd.Email = cfg.Admin
		cmd.Password = cfg.Password
	}

	cfg := loadConfig()
	zc, err := zimbra.Connect(cfg.Host, cfg.Port, cmd.Email, cmd.Password)
	m := log.Default.Msg
	if err != nil {
		m("Login failed. (%s)", err.Error())
		return nil
	}

	defer zc.Close()
	m("Login succeeded for '%s'.", cmd.Email)
	return nil
}
